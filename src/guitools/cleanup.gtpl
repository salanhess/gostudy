<html>
<head>
<title></title>
</head>
<body>
<form action="/cleanup" method="post">
	<!-- 账号名:<input type="text" name="username">
	密码:<input type="password" name="password"> -->
	邮箱:<input type="text" name="email">
	tenant_id:<input type="text" name="tenant_id">
    操作：<select name="operation">
	<option value="recycledeletedvol">回收删除状态硬盘的占用空间</option>
    <option value="delvolume">删除可用状态云盘</option>
    <option value="delattachment">删除挂载状态云盘</option>
    <option value="delsnapshot">删除快照(试用版)</option>
    <option value="delvolid">通过volid删除盘(试用版)</option>
	<option value="checkvolid">通过volid删除盘(试用版)</option>
	<option value="checkquota">检查quota</option>
	<option value="checkjss">检查快照(jss云存储)服务</option>
	<option value="checklogold10">检查zbs相关模块错误日志with10mins(试用版)</option>
	<option value="checklogrange">检查zbs相关模块错误日志accordingRange(试用版)</option>
    </select>
	操作原因:<input type="text" name="note">
	<input type="hidden" name="token" value="{{.}}">
	<input type="submit" value="提交">
</form>
</body>
</html>
