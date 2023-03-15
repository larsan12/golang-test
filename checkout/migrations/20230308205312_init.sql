-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.users
(
    user_id BIGSERIAL PRIMARY KEY,
    name text
);

CREATE TABLE IF NOT EXISTS public.cart_items (
    user_id bigint NOT NULL,
    sku bigint NOT NULL,
    count integer NOT NULL,
    CONSTRAINT cart_items_pk PRIMARY KEY (user_id, sku)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.users;
DROP TABLE IF EXISTS public.cart_items;
-- +goose StatementEnd
