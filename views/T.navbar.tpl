{{define "navbar"}}
<a class="navbar-brand" href="/">后台管理</a>
<div>
	<ul class="nav navbar-nav">
		<li {{if .IsHome}}class="active"{{end}}><a href="/">首页</a></li>
		<li {{if .IsCategory}}class="active"{{end}}><a href="/admin/category/list">分类</a></li>
		<li {{if .IsTopic}}class="active"{{end}}><a href="/admin/article/list">文章</a></li>
	</ul>
</div>

<div class="pull-right">
	<ul class="nav navbar-nav">
		{{if .IsLogin}}
		<li><a href="/login?exit=true">退出登录</a></li>
		{{else}}
		<li><a href="/login">管理员登录</a></li>
		{{end}}
	</ul>
</div>
{{end}}
