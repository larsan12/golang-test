package workerpool

import (
	"context"
	"errors"
)

// описание внешнего интерфейса

type Task[In, Out any] struct {
	Run    func(In) Out
	InArgs In
}

type WorkerPool[In, Out any] interface {
	Execute(context.Context, []Task[In, Out]) ([]Out, error)
	Close()
}

// проверка типов
var _ WorkerPool[any, any] = implemetation[any, any]{}

/**
	по задумке - глобально делаем worker pool и запускаем N воркеров, один общий канал для входящих задач
	каждый пользователь пула вызывает метод Execute с набором задач, в ответ получает массив с результатами
	для каждого запроса Execute создаётся отдельный канал куда воркеры отправляют ответы, после канал закрывается
**/

// добавляем каналы в таску
// outChannel - канал для результатов выполнения, для каждого набора тасок свой канал
// stopChannel - канал на случай непредвиденного закрытия канала outChannel
// stopChannel закрывается прямо перед закрытием outChannel
type TaskWithChannel[In, Out any] struct {
	Task[In, Out]
	outChannel  chan<- Out
	stopChannel <-chan struct{}
}

// реализуем воркера
func worker[In, Out any](ctx context.Context, tasks <-chan TaskWithChannel[In, Out]) {
	// случаем глобальный канал с тасками
	for task := range tasks {
		var result Out

		// проверям не закрыт ли контекст, иначе завершаем воркера
		select {
		case <-ctx.Done():
			return
		default:
		}

		// пробуем выполнить таску
		// если stopChannel закрыт то пропускаем итерацию и не выполняем таску
		// иначе выполняем таску
		select {
		case <-task.stopChannel:
			continue
		default:
			result = task.Run(task.InArgs)
		}

		// если stopChannel не закрыт - отправляем в outChannel результат
		select {
		case <-task.stopChannel:
		default:
			task.outChannel <- result
		}
	}
}

// globalTasksChanel - общий канал с тасками для всех воркеров
// globalStopChannel - канал для остановки, закрываем его перед закрытием globalTasksChanel
type implemetation[In, Out any] struct {
	workersAmount     int
	globalTasksChanel chan TaskWithChannel[In, Out]
	globalStopChannel chan struct{}
}

// создаём один входящий канал для задач и запускаем воркеров
func NewPool[In, Out any](ctx context.Context, workersAmount int) WorkerPool[In, Out] {
	globalTasksChanel := make(chan TaskWithChannel[In, Out], 100)
	globalStopChannel := make(chan struct{})

	for i := 0; i < workersAmount; i++ {
		go func() {
			worker(ctx, globalTasksChanel)
		}()
	}

	return &implemetation[In, Out]{
		workersAmount,
		globalTasksChanel,
		globalStopChannel,
	}
}

// основной метод
// возвращает массив с результатами после выполнения, данные возвращаются в произвольном порядке
// возвращаем массив результатов, не можем вернуть канал с данными поскольку должны создавать и закрывать канал внутри одной сущности
func (p implemetation[In, Out]) Execute(ctx context.Context, tasks []Task[In, Out]) ([]Out, error) {
	// для каждого вызова создаётся канал с результатами
	outChannel := make(chan Out, len(tasks))
	defer close(outChannel)
	stopChannel := make(chan struct{})
	defer close(stopChannel)

	// отправляем задачи
	for _, Task := range tasks {
		// проверяем каналы ctx.Done() и p.globalStopChannel перед отправкой тасок
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-p.globalStopChannel:
			return nil, errors.New("WorkerPool is closed")
		default:
			p.globalTasksChanel <- TaskWithChannel[In, Out]{
				Task,
				outChannel,
				stopChannel,
			}
		}
	}

	// получаем ответы
	result := make([]Out, 0, len(tasks))
	for i := 0; i < len(tasks); i++ {
		// ждём ответы, слушаем каналы ctx.Done() и p.globalStopChannel
		select {
		case <-ctx.Done():
			return result, ctx.Err()
		case <-p.globalStopChannel:
			return result, errors.New("WorkerPool is closed")
		case res := <-outChannel:
			result = append(result, res)
		}
	}

	return result, nil
}

func (p implemetation[In, Out]) Close() {
	close(p.globalStopChannel)
	close(p.globalTasksChanel)
}
