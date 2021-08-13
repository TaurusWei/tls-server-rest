/*
    区块链代理交易信息表 fb_transaction_info
    区块链事件监听块信息表 fb_proxy_block_listen_event

    代理区块事件信息表 fb_proxy_register_block_event
    代理智能合约事件信息表 fb_proxy_register_event

    代理事件待处理信息表 fb_proxy_pending_event
    通道信息表 fb_channel
    网络信息表 fb_network
    order信息表 fb_orderer
    组织信息表 fb_organization
    peer信息表 fb_peer
    成员信息表 ur_user

*/

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


-- 通道表
drop table if exists fb_channel;
create table fb_channel
(
    id serial primary key,
    tenant_id int default 0 not null,
    network_id int default 0,
    name varchar(64),
    desc_ varchar(128),
    creator_org_id int default null,
    creator int,
    creator_name varchar(128) default null,
    create_time bigint default 0
);
create unique index idx_fb_channel on fb_channel(id);

comment on table fb_channel is '通道表';
comment on column fb_channel.id is '通道id';
comment on column fb_channel.tenant_id is '租户id（暂不使用）';
comment on column fb_channel.network_id is '网络区块链id';
comment on column fb_channel.name is '唯一标识';
comment on column fb_channel.desc_ is '描述';
comment on column fb_channel.creator_org_id is '发起组织mspid';
comment on column fb_channel.creator is '创建/申请加入 人员编号';
comment on column fb_channel.creator_name is '创建/申请加入 者名称';
comment on column fb_channel.create_time is '创建/申请加入 时间';

-- 合约相关
drop table if exists fb_contract;
create table fb_contract
(
    id  serial not null primary key,
    tenant_id int default '0' not null,
    network_id int not null,
    org_id int default 0,
	baseorg_mspid varchar(64) default '',
    name varchar(64),
    desc_ varchar(128),
    channel_id int not null,
    channel varchar(256) not null,
    lang smallint default 1,
    status smallint default 0
);
create unique index idx_fb_contract on fb_contract(id);

comment on table fb_contract is '智能合约表';
comment on column fb_contract.id is '智能合约id';
comment on column fb_contract.tenant_id is '租户id（暂不使用）';
comment on column fb_contract.network_id is '网络区块链id';
comment on column fb_contract.org_id is '参与组织id';
comment on column fb_contract.baseorg_mspid is '私链模式，创建网络基础组织 mspid';
comment on column fb_contract.name is '唯一标识';
comment on column fb_contract.desc_ is '描述';
comment on column fb_contract.channel_id is '通道id';
comment on column fb_contract.channel is '通道名称';
comment on column fb_contract.lang is '1-go 2-node.js 3-java';
comment on column fb_contract.status is '1-启用 2-禁用';

-- 网络表
drop table if exists fb_network;
create table fb_network
(
    id bigserial primary key,
    tenant_id int default 0 not null,
    name varchar(64),
    domain varchar(128) default null,
    desc_ varchar(128),
	type smallint default 0,
    creator int,
    creator_name varchar(128),
    create_time bigint not null,
    update_time bigint default 0,
    multi_channel char(1) default '0' ,
    is_inited_local char(1) default '0',
	baseorg_mspid varchar(64) default '',
    status smallint default 0
);
create unique index idx_fb_network on fb_network(id);

comment on table fb_network is '区块链网络表';
comment on column fb_network.id is '区块链网络id';
comment on column fb_network.tenant_id is '租户id（暂不使用）';
comment on column fb_network.name is '区块链网络名称';
comment on column fb_network.domain is '域';
comment on column fb_network.desc_ is '描述';
comment on column fb_network.type is '类型 0-联盟链 1-私有链';
comment on column fb_network.creator is '创建人员/申请加入人员编号';
comment on column fb_network.creator_name is '创建人员/申请加入人员姓名';
comment on column fb_network.create_time is '创建人员/申请加入人员时间';
comment on column fb_network.update_time is '修改时间';
comment on column fb_network.multi_channel is '是否支持多链 0-否 1-是';
comment on column fb_network.is_inited_local is '由本组织创建 0-否 1-是';
comment on column fb_network.baseorg_mspid is '私链模式，创建网络基础组织 mspid';
comment on column fb_network.status is '网络状态 0-未初始化 1-申请加入中 2-正常 11-禁用';

-- order
drop table if exists fb_orderer;
create table fb_orderer
(
  id  serial primary key,
  network_id int default 0,
  name varchar(256) not null,
  ip_address varchar(256) default '',
  port varchar(256) default '',
  tls_cert text,
  tls_key text,
  msp_cert text,
  msp_key text,
  log_level varchar(256),
  status int default 0,
  auth_kafka_tls boolean ,
  kafka_tls_cacert text,
  kafka_tls_keystore text,
  kafka_tls_cert text,
  org_id int not null
);
create unique index idx_fb_orderer on fb_orderer(id);

comment on table fb_orderer is 'orderer表';
comment on column fb_orderer.network_id is '所在网络id，默认0兼容已有api单网络模式';


-- organization
drop table if exists fb_organization;
create table fb_organization
(
  id serial primary key,
  tenant_id int default 0 not null,
  network_id int default 0,
  msp_id varchar(256) not null,
  name varchar(1024) not null,
  org_name varchar(256) default null,
  tls_ca_cert text,
  tls_ca_key text,
  msp_ca_cert text,
  msp_ca_key text
);
create unique index idx_fb_organization on fb_organization(id);

comment on table fb_organization is '区块链组织表';
comment on column fb_organization.tenant_id is '租户id（暂不使用）';
comment on column fb_organization.network_id is '所在网络id，默认0兼容已有api单网络模式';
comment on column fb_organization.name is '组织domain';
comment on column fb_organization.org_name is '组织名称';


-- peer
drop table if exists fb_peer;
create table fb_peer
(
  id serial primary key,
  network_id int default 0,
  name varchar(256) not null,
  ip_address varchar(256) default '',
  port varchar(256) default '',
  cc_port varchar(256) default '',
  event_port varchar(256) default '',
  tls_cert text,
  tls_key text,
  msp_cert text,
  msp_key text,
  log_level varchar(256),
  org_id int not null,
  baseorg_msp varchar(256) default '',
  boot_up varchar(256),
  db_type varchar(256) default '',
  couchdb_id int default null,
  careate_time timestamp
);
create unique index idx_fb_peer on fb_peer(id);

comment on table fb_peer is 'peer表';
comment on column fb_peer.network_id is '所在网络id，默认0兼容已有api单网络模式';



-- 成员表
drop table if exists ur_user;
create table ur_user
(
  id bigserial primary key,
  name varchar(256) not null,
  password varchar(256) default null,
  tls_cert text,
  tls_key text,
  msp_cert text,
  msp_key text,
  network_id int default 0,
  tenant_id int default 0 not null,
  member_name varchar(64),
  member_desc varchar(128),
  org_id int not null,
  baseorg_mspid varchar(64) default '',
  owner_name varchar(64) default null,
  status smallint default 1,
  creator int,
  creator_name varchar(64) default null ,
  create_time bigint default 0,
  expire_time bigint default 0,
  email varchar(64) default null,
  mobile varchar(64) default null,
  purpose smallint default 0
);
create unique index idx_ur_user on ur_user(id);

comment on table ur_user is '成员表';
comment on column ur_user.password is 'sha256密码哈希';
comment on column ur_user.network_id is '所在网络id，默认0兼容已有api单网络模式';
comment on column ur_user.tenant_id is '租户id（暂不使用）';
comment on column ur_user.member_name is '成员名称';
comment on column ur_user.member_desc is '成员描述';
comment on column ur_user.org_id is '所属组织机构id';
comment on column ur_user.baseorg_mspid is '私链模式，创建网络基础组织 mspid';
comment on column ur_user.owner_name is '负责人（成员姓名）';
comment on column ur_user.status is '状态 1-正常(默认) 2-停用';
comment on column ur_user.creator is '创建人员编号';
comment on column ur_user.creator_name is '创建人姓名';
comment on column ur_user.create_time is '创建时间';
comment on column ur_user.expire_time is '证书过期时间';
comment on column ur_user.email is '负责人邮箱';
comment on column ur_user.mobile is '负责人手机号';
comment on column ur_user.purpose is '用途 1-业务应用（默认） 2-业务用户 3-组织机构';
