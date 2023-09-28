INSERT INTO hh_admin.sys_menu (menu_name,parent_id,order_num,`path`,component,query,is_frame,is_cache,menu_type,visible,status,perms,icon,create_by,update_by,remark,create_time,update_time) VALUES
	 ('权限管理',0,1,'/api/v1/system/auth','','',1,1,1,1,1,'/api/v1/system/auth:GET','#','1','1','','2023-09-07 23:34:20','2023-09-07 23:34:23'),
	 ('系统监控',0,1,'/api/v1/system/monitor','','',1,1,1,1,1,'/api/v1/system/monitor:GET','#','1','1','','2023-09-07 23:36:14','2023-09-07 23:34:27'),
	 ('菜单管理',1,2,'/api/v1/system/auth/menu/manage','','',1,1,2,1,1,'/api/v1/system/auth/menu/manage:GET','#','1','1','','2023-09-07 23:36:16','2023-09-07 23:36:18'),
	 ('角色管理',1,2,'/api/v1/system/auth/role/manage','','',1,1,2,1,1,'/api/v1/system/auth/role/manage:GET','#','1','1','','2023-09-07 23:37:13','2023-09-07 23:37:15'),
	 ('部门管理',1,2,'/api/v1/system/auth/dept/manage','','',1,1,2,1,1,'/api/v1/system/auth/dept/manage:GET','#','1','1','','2023-09-07 23:37:13','2023-09-07 23:37:15'),
	 ('岗位管理',1,2,'/api/v1/system/auth/post/manage','','',1,1,2,1,1,'/api/v1/system/auth/post/manage:GET','#','1','1','','2023-09-07 23:37:13','2023-09-07 23:37:15'),
	 ('用户管理',1,2,'/api/v1/system/auth/user/manage','','',1,1,2,1,1,'/api/v1/system/auth/user/manage:GET','#','1','1','','2023-09-07 23:37:13','2023-09-07 23:37:15'),
	 ('服务监控',2,2,'/api/v1/system/monitor/server/manage','','',1,1,2,1,1,'/api/v1/system/monitor/server/manage:GET','#','1','1','','2023-09-07 23:37:13','2023-09-07 23:37:15'),
	 ('登录日志',2,2,'/api/v1/system/monitor/login/manage','','',1,1,2,1,1,'/api/v1/system/monitor/login/manage:GET','#','1','1','','2023-09-07 23:37:13','2023-09-07 23:37:15'),
	 ('操作日志',2,2,'/api/v1/system/monitor/oper/manage','','',1,1,2,1,1,'/api/v1/system/monitor/oper/manage:GET','#','1','1','','2023-09-07 23:37:13','2023-09-07 23:37:15');
INSERT INTO hh_admin.sys_menu (menu_name,parent_id,order_num,`path`,component,query,is_frame,is_cache,menu_type,visible,status,perms,icon,create_by,update_by,remark,create_time,update_time) VALUES
	 ('在线用户',2,2,'/api/v1/system/monitor/online/manage','','',1,1,2,1,1,'/api/v1/system/monitor/online/manage:GET','#','1','1','','2023-09-07 23:37:13','2023-09-07 23:37:15'),
	 ('角色列表',4,1,'/api/v1/system/auth/role/list','','',1,1,3,1,1,'/api/v1/system/auth/role/list:GET','#','1','1','','2023-09-10 01:18:52','2023-09-10 01:18:54'),
	 ('角色新增',4,2,'/api/v1/system/auth/role/add','','',1,1,3,1,1,'/api/v1/system/auth/role/add:POST','#','1','1','','2023-09-10 01:18:52','2023-09-10 01:18:54'),
	 ('角色修改',4,3,'/api/v1/system/auth/role/update','','',1,1,3,1,1,'/api/v1/system/auth/role/update:POST','#','1','1','','2023-09-10 01:18:52','2023-09-10 01:18:54'),
	 ('角色删除',4,4,'/api/v1/system/auth/role/delete','','',1,1,3,1,1,'/api/v1/system/auth/role/delete:POST','#','1','1','','2023-09-10 01:18:52','2023-09-10 01:18:54'),
	 ('部门列表',5,1,'/api/v1/system/auth/dept/list','','',1,1,3,1,1,'/api/v1/system/auth/dept/list:GET','#','1','1','','2023-09-10 01:18:52','2023-09-10 01:18:54'),
	 ('部门新增',5,2,'/api/v1/system/auth/dept/add','','',1,1,3,1,1,'/api/v1/system/auth/dept/add:POST','#','1','1','','2023-09-10 01:18:52','2023-09-10 01:18:54'),
	 ('部门修改',5,3,'/api/v1/system/auth/dept/update','','',1,1,3,1,1,'/api/v1/system/auth/dept/update:POST','#','1','1','','2023-09-10 01:18:52','2023-09-10 01:18:54'),
	 ('部门删除',5,4,'/api/v1/system/auth/dept/delete','','',1,1,3,1,1,'/api/v1/system/auth/dept/delete:POST','#','1','1','','2023-09-10 01:18:52','2023-09-10 01:18:54'),
	 ('岗位查询',6,1,'/api/v1/system/auth/post/list','','',1,1,3,1,1,'/api/v1/system/auth/post/list:GET','#','1','1','','2023-09-10 01:18:52','2023-09-10 01:18:54');
INSERT INTO hh_admin.sys_menu (menu_name,parent_id,order_num,`path`,component,query,is_frame,is_cache,menu_type,visible,status,perms,icon,create_by,update_by,remark,create_time,update_time) VALUES
	 ('岗位新增',6,2,'/api/v1/system/auth/post/add','','',1,1,3,1,1,'/api/v1/system/auth/post/add:POST','#','1','1','','2023-09-10 01:18:52','2023-09-10 01:18:54'),
	 ('岗位修改',6,3,'/api/v1/system/auth/post/update','','',1,1,3,1,1,'/api/v1/system/auth/post/update:POST','#','1','1','','2023-09-10 01:18:52','2023-09-10 01:18:54'),
	 ('岗位删除',6,4,'/api/v1/system/auth/post/delete','','',1,1,3,1,1,'/api/v1/system/auth/post/delete:POST','#','1','1','','2023-09-10 01:18:52','2023-09-10 01:18:54'),
	 ('用户查询',5,1,'/api/v1/system/auth/user/list','','',1,1,3,1,1,'/api/v1/system/auth/user/list:GET','#','1','1','','2023-09-10 01:18:52','2023-09-10 01:18:54'),
	 ('用户新增',5,2,'/api/v1/system/auth/user/add','','',1,1,3,1,1,'/api/v1/system/auth/user/add:POST','#','1','1','','2023-09-10 01:18:52','2023-09-10 01:18:54'),
	 ('用户修改',5,3,'/api/v1/system/auth/user/update','','',1,1,3,1,1,'/api/v1/system/auth/user/update:POST','#','1','1','','2023-09-10 01:18:52','2023-09-10 01:18:54'),
	 ('用户删除',5,4,'/api/v1/system/auth/user/delete','','',1,1,3,1,1,'/api/v1/system/auth/user/delete:POST','#','1','1','','2023-09-10 01:18:52','2023-09-10 01:18:54'),
	 ('菜单查询',3,1,'/api/v1/system/auth/menu/list','','',1,1,3,1,1,'/api/v1/system/auth/menu/list:GET','#','1','1','','2023-09-10 01:18:52','2023-09-10 01:18:54'),
	 ('菜单新增',3,2,'/api/v1/system/auth/menu/add','','',1,1,3,1,1,'/api/v1/system/auth/menu/add:POST','#','1','1','','2023-09-10 01:18:52','2023-09-10 01:18:54'),
	 ('菜单修改',3,3,'/api/v1/system/auth/menu/update','','',1,1,3,1,1,'/api/v1/system/auth/menu/update:POST','#','1','1','','2023-09-10 01:18:52','2023-09-10 01:18:54');
INSERT INTO hh_admin.sys_menu (menu_name,parent_id,order_num,`path`,component,query,is_frame,is_cache,menu_type,visible,status,perms,icon,create_by,update_by,remark,create_time,update_time) VALUES
	 ('菜单删除',3,4,'/api/v1/system/auth/menu/delete','','',1,1,3,1,1,'/api/v1/system/auth/menu/delete:POST','#','1','1','','2023-09-10 01:18:52','2023-09-10 01:18:54');
