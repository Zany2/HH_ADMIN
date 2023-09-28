INSERT INTO hh_admin.casbin_rule (p_type,v0,v1,v2,v3,v4,v5) VALUES
	 ('g','test1','admin_role','','','',''),
	 ('p','admin_post','/api/v1/system/auth','GET','','',''),
	 ('p','admin_post','/api/v1/system/auth/post/add','POST','','',''),
	 ('p','admin_post','/api/v1/system/auth/post/delete','POST','','',''),
	 ('p','admin_post','/api/v1/system/auth/post/list','GET','','',''),
	 ('p','admin_post','/api/v1/system/auth/post/manage','GET','','',''),
	 ('p','admin_post','/api/v1/system/auth/post/update','POST','','',''),
	 ('p','admin_role','/api/v1/system/auth','GET','','',''),
	 ('p','admin_role','/api/v1/system/auth/role/add','POST','','',''),
	 ('p','admin_role','/api/v1/system/auth/role/delete','POST','','','');
INSERT INTO hh_admin.casbin_rule (p_type,v0,v1,v2,v3,v4,v5) VALUES
	 ('p','admin_role','/api/v1/system/auth/role/list','GET','','',''),
	 ('p','admin_role','/api/v1/system/auth/role/manage','GET','','',''),
	 ('p','admin_role','/api/v1/system/auth/role/update','POST','','','');
