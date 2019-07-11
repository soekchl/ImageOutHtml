package main

// format src
const imgModel = `
<img src="%s" style="vertical-align:bottom">
<br/>
`

const htmlStyle = `
		<style type="text/css">
			* { margin: 0; padding: 0; } .all { position: fixed; /*窗口固定定位*/ right:
			0px; top: 70%; width: 60px; margin-top: -90px; z-index: 999; } .all ul
			{ list-style: none; } .all ul li { width: 55px; height: 30px; line-height:
			30px; position: relative; /*相对*/ z-index: 2; } .all ul li a { position:
			absolute; /*绝对*/ top: 0px; left: 0px; display: block; color: #FFFFFF; width:
			60px; height: 30px; line-height: 30px; z-index: 2; text-decoration: none;
			/*下划线样式*/ -webkit-transition: all 0.6s; -ms-transition: all 0.6s; -moz-transition:
			all 0.6s; } .all ul li a img { width: 30px; position: absolute; top: 15px;
			left: 15px; z-index: 2; } li a { background-image: -webkit-linear-gradient(left,
			#f60, #ffb443); background-image: -moz-linear-gradient(left, #f60, #ffb443);
			background-image: -ms-linear-gradient(left, #f60, #ffb443); } li.wuyou-contact
			a { background-image: -webkit-linear-gradient(left, #00b7ee, #55d8ff);
			background-image: -moz-linear-gradient(left, #00b7ee, #55d8ff); background-image:
			-ms-linear-gradient(left, #00b7ee, #55d8ff); } li.wuyou-top a { background-image:
			-webkit-linear-gradient(left, #333, #666); background-image: -moz-linear-gradient(left,
			#333, #666); background-image: -ms-linear-gradient(left, #333, #666); }
			img { width:99%; height:auto; }
		</style>
`

// format     preAddress,NextAddress,images
const htmlModel = `
<!DOCTYPE html>
<html lang="en">
	<head>
		<meta name="apple-mobile-web-app-capable" content="yes" />
		<meta name="apple-mobile-web-app-status-bar-style" content="black" />
		<meta name="viewport" content="width=device-width,minimum-scale=1.0,maximum-scale=1.0,user-scalable=no"/>
		<meta charset="UTF-8">
		<title>%s만화</title>
	%s
	</head>
	<body>
		<div class="all">
			<ul>
				<li><a href="/">首 页</a></li>
				<li class="wuyou-contact"><a href="%s">上一章</a></li>
				<li class="wuyou-top"><a href="%s">下一章</a></li>
			</ul>
		</div>
		<div align="center">
			%s
		</div>
	</body>
</html>
`
