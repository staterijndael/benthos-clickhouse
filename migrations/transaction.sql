create table transaction
(
    transaction_id Nullable(String),
    from_user_id Nullable(Int64),
    to_user_id Nullable(Int64),
    currency_symbol Nullable(String),
    amount Nullable(String)
)
    engine = Memory;

