CREATE DATABASE `tudou_list` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

show databases;
show table status;

use tudou_list;

-- target基本信息表
CREATE TABLE td_f_target(
    id INT NOT NULL AUTO_INCREMENT  COMMENT 'target标识 target标识' ,
    detail VARCHAR(1024) NOT NULL   COMMENT 'target内容 target内容' ,
    feedback VARCHAR(1024)    COMMENT 'target最终完成说明 target最终完成说明' ,
    created_by VARCHAR(32) NOT NULL   COMMENT 'target创建人 target创建人' ,
    created_at DATETIME NOT NULL   COMMENT 'target开始时间 target创建时间' ,
    state INT NOT NULL   COMMENT 'target状态 target状态' ,
    done_at DATETIME    COMMENT 'target结束时间 target结束时间' ,
    deadline DATETIME    COMMENT 'target ddl target ddl' ,
    gap_order INT NOT NULL   COMMENT 'target顺序 target跳跃排序序号' ,
    PRIMARY KEY (id)
) COMMENT = 'target基本信息';

INSERT INTO `tudou_list`.`td_f_target`
(`tgt_id`,
`tgt_detail`,
`tgt_final_note`,
`tgt_created_by`,
`tgt_start_date`,
`tgt_state`,
`tgt_end_date`,
`tgt_order`)
VALUES
(test,
test,
test,
test,
t,
1,
<{tgt_end_date: }>,
1);

select max(gap_order) from td_f_target;
select * from td_f_target;