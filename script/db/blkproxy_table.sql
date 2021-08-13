-- ========== 区块链代理交易信息表 by wsp===========
DROP TABLE IF EXISTS fb_transaction_info;
CREATE TABLE fb_transaction_info (
	id  bigserial primary key,
	network_id int not null,
	username varchar(256) not null,
	channel_name varchar(64) not null,
	contract_name varchar(64) not null,
	call_type varchar(1) not null,
	fcn varchar(64) not null,
	org_msp_id varchar(256) not null,
	business_no varchar(32) not null,
	business_type varchar(64) default '其它',
	req_type varchar(1) default '0',
	tx_id varchar(256) default '',
	status smallint default 0,
	create_time	timestamp default current_timestamp,
	update_time timestamp default current_timestamp
);
create unique index idx_fb_trans_info on fb_transaction_info(business_no);

comment on table fb_transaction_info is '交易信息表';
comment on column fb_transaction_info.network_id is '网络id';
comment on column fb_transaction_info.userName is '用户名';
comment on column fb_transaction_info.channel_name is '通道名';
comment on column fb_transaction_info.contract_name is '合约名';
comment on column fb_transaction_info.call_type is '调用类型，0：query,1: invoke';
comment on column fb_transaction_info.fcn is '调用方法';
comment on column fb_transaction_info.org_msp_id is '组织mspId';
comment on column fb_transaction_info.business_no is '业务流水号';
comment on column fb_transaction_info.business_type is '业务类型';
comment on column fb_transaction_info.req_type is '请求类型，0：同步方式(默认) 1：异步方式';
comment on column fb_transaction_info.tx_id is '交易id';
comment on column fb_transaction_info.status is '交易状态,默认0，0: 未处理，1: 处理中，2: 处理完成，3: 处理出错';
comment on column fb_transaction_info.create_time is '创建时间，默认当前时间';
comment on column fb_transaction_info.update_time is '更新时间，默认当前时间';


-- ========== 区块链代理事件注册 by wbs===========
DROP TABLE IF EXISTS fb_proxy_block_listen_event;
create table fb_proxy_block_listen_event(
   id serial not null primary key,
   network_id  int not null,
   channel_id varchar(256),
   block_number int,
   tran_number int,
   receive_time timestamp,
   peer_id int,
   curr_hash varchar(256)
);
create unique index idx_fb_proxy_block_listen_event on fb_proxy_block_listen_event(id);
comment on table fb_proxy_block_listen_event is '区块链事件监听块信息表';
comment on column fb_proxy_block_listen_event.network_id is '所在网络id，默认0兼容已有api单网络模式';
comment on column fb_proxy_block_listen_event.channel_id is '所在通道';
comment on column fb_proxy_block_listen_event.block_number is '区块编号';
comment on column fb_proxy_block_listen_event.tran_number is '交易编号';
comment on column fb_proxy_block_listen_event.receive_time is '接收时间';
comment on column fb_proxy_block_listen_event.peer_id is 'peer节点';
comment on column fb_proxy_block_listen_event.curr_hash is '区块当前hash';

DROP TABLE IF EXISTS fb_proxy_register_block_event;
create table fb_proxy_register_block_event(
   id bigserial not null primary key,
   user_name varchar(256),
   msp_id varchar(256),
   network_id int not null,
   channel_id varchar(256),
   callback_url varchar(1024),
   serial_number varchar(32),
   create_time  timestamp,
   event_type int,
   req_type int
);
create unique index idx_fb_proxy_register_block_event on fb_proxy_register_block_event(id);
comment on table fb_proxy_register_block_event is '代理区块事件信息表';
comment on column fb_proxy_register_block_event.user_name is '用户名';
comment on column fb_proxy_register_block_event.msp_id is '组织成员Id';
comment on column fb_proxy_register_block_event.network_id is '所在网络id，默认0兼容已有api单网络模式';
comment on column fb_proxy_register_block_event.channel_id is '通道';
comment on column fb_proxy_register_block_event.callback_url is '回调路径';
comment on column fb_proxy_register_block_event.serial_number is '业务流水号';
comment on column fb_proxy_register_block_event.create_time is '创建时间';
comment on column fb_proxy_register_block_event.event_type is '事件类型';
comment on column fb_proxy_register_block_event.req_type is '请求类型';

DROP TABLE IF EXISTS fb_proxy_register_event;
create table fb_proxy_register_event(
   id bigserial not null primary key,
   user_name varchar(256),
   msp_id varchar(256),
   network_id int not null,
   channel_id varchar(256),
   cc_name varchar(256),
   event_filter varchar(256),
   callback_url varchar(1024),
   serial_number varchar(32),
   create_time  timestamp,
   event_type int,
   req_type int
);
create unique index idx_fb_proxy_register_event on fb_proxy_register_event(id);
comment on table fb_proxy_register_event is '代理智能合约事件信息表';
comment on column fb_proxy_register_event.user_name is '用户名';
comment on column fb_proxy_register_event.msp_id is '组织成员Id';
comment on column fb_proxy_register_event.network_id is '所在网络id，默认0兼容已有api单网络模式';
comment on column fb_proxy_register_event.channel_id is '通道';
comment on column fb_proxy_register_event.cc_name is '合约名称';
comment on column fb_proxy_register_event.event_filter is '事件名称';
comment on column fb_proxy_register_event.callback_url is '回调路径';
comment on column fb_proxy_register_event.serial_number is '业务流水号';
comment on column fb_proxy_register_event.create_time is '创建时间';
comment on column fb_proxy_register_event.event_type is '事件类型';
comment on column fb_proxy_register_event.req_type is '请求类型';

DROP TABLE IF EXISTS fb_proxy_pending_event;
create table fb_proxy_pending_event(
   id bigserial not null primary key,
   user_name varchar(256),
   msp_id varchar(256),
   network_id int not null,
   callback_url varchar(1024),
   callback_data text,
   serial_number varchar(32),
   event_type int,
   req_type int
);
create unique index idx_fb_proxy_pending_event on fb_proxy_pending_event(id);
comment on table fb_proxy_pending_event is '代理事件待处理信息表';
comment on column fb_proxy_pending_event.user_name is '用户名';
comment on column fb_proxy_pending_event.msp_id is '组织成员Id';
comment on column fb_proxy_pending_event.network_id is '所在网络id，默认0兼容已有api单网络模式';
comment on column fb_proxy_pending_event.callback_url is '回调路径';
comment on column fb_proxy_pending_event.callback_data is '回调数据';
comment on column fb_proxy_pending_event.serial_number is '业务流水号';
comment on column fb_proxy_pending_event.event_type is '事件类型';
comment on column fb_proxy_pending_event.req_type is '请求类型';

