<html>
<script>
alert("My First JavaScript");
</script>
<head>
<title></title>
</head>
<body>
<form action="/display" method="post">
	<!-- 账号名:<input type="text" name="username">
	密码:<input type="password" name="password"> -->
	邮箱:<input type="text" name="email">
	tenant_id:<input type="text" name="tenant_id">
    操作：<select name="operation">
    <option value="delvolume">删除可用状态云盘</option>
    <option value="delattachment">删除挂载状态云盘</option>
    <option value="delsnapshot">删除快照(试用版)</option>
	<option value="checkquota">检查quota</option>
    </select>
	操作原因:<input type="text" name="note">
	<input type="hidden" name="token" value="{{.}}">
	<input type="submit" value="提交">
</form>
</body>
</html>
