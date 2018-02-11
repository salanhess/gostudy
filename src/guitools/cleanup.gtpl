<html>
<head>
<title></title>
</head>
<body>
<table>
<tr>
<td>
<form action="/cleanup" method="post">
	账号名:<input type="text" name="username">
	密码:<input type="password" name="password">
	邮箱:<input type="text" name="email">
	tenant_id:<input type="text" name="tenant_id">
    操作：<select name="operation">
    <option value="delvolume">删除可用状态云盘</option>
    <option value="delattachment">删除挂载状态云盘</option>
    <option value="delsnapshot">删除快照(试用版)</option>
    </select>
	操作原因:<input type="text" name="note">
	<input type="submit" value="提交">
</form>
</td>
</tr>
</body>
</html>
