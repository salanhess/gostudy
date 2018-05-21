<html>
<head>
<title></title>
</head>
<body>
<form action="/build" method="post">
	<!-- 账号名:<input type="text" name="username">
	密码:<input type="password" name="password"> -->
	邮箱:<input type="text" name="email">
    操作：<select name="operation">
	<option value="build_jenkins">调用Jenkins编译</option>
    </select>
	zbs_server_branch:<input type="text" name="zbs_server_branch">
	zbs_server_commitid:<input type="text" name="zbs_server_commitid">		
	zbs_scheduler_branch:<input type="text" name="zbs_scheduler_branch">	
	zbs_scheduler_commitid:<input type="text" name="zbs_scheduler_commitid">
	zbs_common_branch:<input type="text" name="zbs_common_branch">
	zbs_common_commitid:<input type="text" name="zbs_common_commitid">		
	zbs_gateway_branch:<input type="text" name="zbs_gateway_branch">	
	zbs_gateway_commitid:<input type="text" name="zbs_gateway_commitid">	
	zbs_worker_branch:<input type="text" name="zbs_worker_branch">
	zbs_worker_commitid:<input type="text" name="zbs_worker_commitid">		
	zbs_storage_branch:<input type="text" name="zbs_storage_branch">	
	zbs_storage_commitid:<input type="text" name="zbs_storage_commitid">
	zbs_client_branch:<input type="text" name="zbs_client_branch">
	编译模块:<input type="text" name="module">		
	编译tag:<input type="text" name="tag">		
	<input type="hidden" name="token" value="{{.}}">
	<input type="submit" value="提交">
</form>
</body>
</html>
