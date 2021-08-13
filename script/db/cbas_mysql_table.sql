# 区块链代理交易信息表
DROP TABLE IF EXISTS transaction_info;
CREATE TABLE transaction_info (
	id INT AUTO_INCREMENT PRIMARY KEY,
	network_id INT NOT NULL COMMENT '网络id',
	user_name CHAR(255) NOT NULL COMMENT '用户名',
	channel_name CHAR(64) NOT NULL COMMENT '通道名',
	contract_name CHAR(64) NOT NULL COMMENT '合约名',
	call_type CHAR(4) NOT NULL COMMENT '调用类型，INVK：调用, QURY: 查询',
	fcn CHAR(64) NOT NULL COMMENT '调用方法',
	msp_id CHAR(255) NOT NULL COMMENT '组织mspId',
	biz_num CHAR(32) NOT NULL COMMENT '业务流水号',
	biz_type CHAR(64) DEFAULT '0' COMMENT '业务类型',
	req_type CHAR(4) DEFAULT 'SYNC' COMMENT '请求类型，SYNC：同步方式(默认) ASYN：异步方式',
	tx_id CHAR(255) DEFAULT '' COMMENT '交易id',
	tx_stat CHAR(4) COMMENT '交易状态,，HDLU: 未处理，HDLD: 处理完成，HDLF: 处理失败，HDLG: 处理中',
	crt_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间，默认当前时间',
	updt_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间，默认当前时间'
)
engine=innodb default charset=utf8 comment='区块链代理交易信息表';


# 代理注册块事件表
DROP TABLE IF EXISTS proxy_register_block_event;
CREATE TABLE proxy_register_block_event(
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    user_name CHAR(255) COMMENT '用户名',
    msp_id CHAR(255) COMMENT '组织成员Id',
    network_id INT NOT NULL COMMENT '所在网络id，默认0兼容已有api单网络模式',
    channel_id CHAR(255) COMMENT '通道',
    callback_url VARCHAR(1024) COMMENT '回调路径',
    biz_num CHAR(32) COMMENT '业务流水号',
    crt_time  TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    event_type CHAR(4) COMMENT '事件类型 EVTC:链码事件, EVTB: 块事件',
    req_type CHAR(4) COMMENT '请求类型 ASYN：异步， SYNC: 同步'
)
engine=innodb default charset=utf8 comment='代理注册块事件表';

# 代理监听块事件表
DROP TABLE IF EXISTS proxy_block_listen_event;
CREATE TABLE proxy_block_listen_event(
    id CHAR(255) NOT NULL PRIMARY KEY,
    network_id  INT NOT NULL COMMENT '所在网络id，默认0兼容已有api单网络模式',
    channel_id CHAR(255) COMMENT '所在通道',
    block_num BIGINT COMMENT '区块编号',
    tx_cnt INT COMMENT '交易数量',
    recv_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '接收时间',
    peer_id CHAR(255) COMMENT 'peer节点',
    curr_hash CHAR(255) COMMENT '区块当前hash'
)
engine=innodb default charset=utf8 comment='代理监听块事件表';

# 代理注册智能合约事件信息表
DROP TABLE IF EXISTS proxy_register_cc_event;
CREATE TABLE proxy_register_cc_event(
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    user_name CHAR(255) COMMENT '用户名',
    msp_id CHAR(255) COMMENT '组织成员Id',
    network_id INT NOT NULL COMMENT '所在网络id，默认0兼容已有api单网络模式',
    channel_id CHAR(255) COMMENT '通道',
    contract_name CHAR(255) COMMENT '合约名称',
    event_name CHAR(255) COMMENT '事件名称',
    callback_url VARCHAR(1024) COMMENT '回调路径',
    biz_num CHAR(32) COMMENT '业务流水号',
    crt_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    event_type CHAR(4) COMMENT '事件类型 EVTC:链码事件, EVTB: 块事件',
    req_type CHAR(4) COMMENT '请求类型 ASYN：异步， SYNC: 同步'
)
engine=innodb default charset=utf8 comment='代理注册智能合约事件信息表';


# 代理监听智能合约事件信息表
DROP TABLE IF EXISTS proxy_cc_listen_event;
CREATE TABLE proxy_cc_listen_event(
    id CHAR(255) NOT NULL PRIMARY KEY,
    user_name CHAR(255) COMMENT '用户名',
    msp_id CHAR(255) COMMENT '组织成员Id',
    network_id INT NOT NULL COMMENT '所在网络id，默认0兼容已有api单网络模式',
    callback_url VARCHAR(1024) COMMENT '回调路径',
    callback_data BLOB COMMENT '回调数据',
    biz_num CHAR(32) COMMENT '业务流水号',
    event_type CHAR(4) COMMENT '事件类型',
    req_type CHAR(4) COMMENT '请求类型'
)
engine=innodb default charset=utf8 comment='代理监听智能合约事件信息表';


# 监听报文转发信息表
DROP TABLE IF EXISTS msg_forward;
CREATE TABLE msg_forward(
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    listen_cc_id CHAR(255) NOT NULL COMMENT '监听链码信息id',
	blk_id BIGINT COMMENT '所在区块号',
    recv_bank CHAR(255) NOT NULL COMMENT '接受行',
    msg BLOB COMMENT '报文',
    stat CHAR(4) DEFAULT 0 COMMENT '推送状态，默认0，推送失败，1，推送成功',
	target CHAR(255) DEFAULT '' COMMENT '推送目标',
    resp VARCHAR(1024) DEFAULT '' COMMENT '推送回执',
	tm TIMESTAMP(6) COMMENT '操作时间'
#     times INT COMMENT '推送次数'
)
engine=innodb default charset=utf8 comment='监听报文转发信息表';

CREATE INDEX IX1_STAT_MSG_FORWARD ON msg_forward(stat);
CREATE INDEX IX2_ID_STAT_MSG_FORWARD ON msg_forward(id, stat);
CREATE INDEX IX3_BLK_ID_MSG_FORWARD ON msg_forward(blk_id);


