SELECT active_period, quiet_period, active_num, quiet_num
  FROM wx_group
 WHERE group_id == %%GroupID string%%;
