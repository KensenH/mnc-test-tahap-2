CREATE TABLE users (
  user_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  first_name VARCHAR(50) NOT NULL,
  last_name  VARCHAR(50),
  phone_number VARCHAR(50) unique,
  address VARCHAR(50), 
  pin VARCHAR(6) not null,
  balance decimal default 0,
  created_date TIMESTAMP default now()
);

create table top_ups (
	top_up_id UUID PRIMARY key DEFAULT uuid_generate_v4(),
	user_id UUID references users(user_id),
	amount_top_up decimal not null default 0,
	status varchar(10) not null default 'pending',
	created_date timestamp default now()
);

create table payments (
	payment_id UUID PRIMARY key DEFAULT uuid_generate_v4(),
	user_id UUID references users(user_id),
	amount decimal not null default 0,
	status varchar(10) not null default 'pending',
	remarks text not null,
	created_date timestamp default now()
);

create table transfers (
	transfer_id UUID PRIMARY key DEFAULT uuid_generate_v4(),
	user_id UUID references users(user_id),
	target_user UUID references users(user_id),
	amount decimal not null default 0,
	remarks text,
	created_date timestamp default now()
);

create table mutations (
	id UUID PRIMARY key DEFAULT uuid_generate_v4(),
	user_id UUID references users(user_id),
	amount decimal not null default 0,
	balance_before decimal not null default 0,
	balance_after decimal not null default 0,
	transaction_type varchar(10) not null,
	status varchar(10),
	created_date timestamp default now()
)