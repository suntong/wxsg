-- Exported from QuickDBD: https://www.quickdatabasediagrams.com/
-- NOTE! If you have used non-SQL datatypes in your design, you will have to change these here.


-- WX user
CREATE TABLE user (
    user_id  INTEGER  NOT NULL,
    -- user name
    -- WX 昵称（不是群昵称）
    name string   NOT NULL,
    CONSTRAINT pk_user PRIMARY KEY (
        user_id
     )
);

-- WX 群
CREATE TABLE wx_group (
    group_id  INTEGER  NOT NULL,
    -- Group name
    name string   NOT NULL,
    -- Group comment
    comment string   NULL,
    -- Day range for reporting of being active
    active_period int  DEFAULT 15 NOT NULL,
    -- Day range for reporting of being sleep
    quiet_period int  DEFAULT 30 NOT NULL,
    -- Number of active users to report
    active_num int  DEFAULT 10 NOT NULL,
    -- Number of quiet users to report
    quiet_num int  DEFAULT 10 NOT NULL,
    CONSTRAINT pk_wx_group PRIMARY KEY (
        group_id
     )
);

CREATE TABLE activity_type (
    activity_type_id  INTEGER  NOT NULL,
    name string   NOT NULL,
    -- scale factor for the activity_type
    scale int   NOT NULL,
    -- 贡献	发言总字数
    -- 爱心	发言后又撤回重新修改的总字数
    -- 分享	除纯文字以外，任何其他的分享
    -- 热闹	贴表情包图
    -- .
    -- .
    -- .
    -- .
    -- .
    -- .
    -- .
    comment string   NULL,
    CONSTRAINT pk_activity_type PRIMARY KEY (
        activity_type_id
     )
);

CREATE TABLE activity_log (
    activity_log_id  INTEGER  NOT NULL,
    user_id int   NOT NULL,
    group_id int   NOT NULL,
    activity_type_id int   NOT NULL,
    activity_date dateTime   NOT NULL,
    stat int   NOT NULL,
    CONSTRAINT pk_activity_log PRIMARY KEY (
        activity_log_id
     )
);

-- user statistic (View structure)
CREATE TABLE user_stat_ (
    user_id int   NOT NULL,
    group_id int   NOT NULL,
    -- 贡献指数	发言总字数累计
    word_count int   NOT NULL,
    -- 爱心指数	发言后又撤回重新修改的总字数累计
    correction_count int   NOT NULL,
    -- 分享指数	除纯文字以外，任何其他分享的次数累计
    share_count int   NOT NULL,
    -- 热闹指数	贴表情包图的次数累计（表情包不统计在分享指数之中）
    meme_count int   NOT NULL
);






CREATE INDEX idx_user_name
ON user (name);

CREATE INDEX idx_wx_group_name
ON wx_group (name);

