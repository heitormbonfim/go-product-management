# How to run

From the root directory

```bash
go run cmd/main.go
```

## Docker

```bash
sudo docker compose up -d go_db
```

## Tables Creation

1. Create product table

```sql
CREATE TABLE product (
  id SERIAL PRIMARY KEY,
  product_name VARCHAR(255) NOT NULL,
  price NUMERIC(10, 2)
)
```

2. Insert product into product table

```sql
INSERT INTO product (
	product_name,
	price
)
VALUES (
	'Weed',
	4.20
)
```
