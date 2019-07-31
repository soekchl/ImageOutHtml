package main

// format src
const imgModel = `
<img data-src="%s" name="%s" style="vertical-align:bottom">
<br/>
`

const htmlStyle = `
 <style type="text/css">
img { width:100%; height:auto; }
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
		<div id="imgList" align="center">
			%s
		</div>
	</body>
<script>
  let lastIndex = 0
  const unit = 8
  let srcTitle = document.title
  init()

  window.addEventListener('scroll', function () {
    let index = getShowImgIndex()
    let start = index - unit
    let stop = index + unit

    if (start < 0) {
      start = 0
    }
    if (stop > imgList.children.length) {
      stop = imgList.children.length
    }

    // 遍历所有的img标签
    for (let i = start; i < stop; i += 2) {
      loadImg(imgList.children[i])
    }
    // console.log(start, stop)
    for (let i = start - 2; i > 0 && setEmptyImg(i); i -= 2) { // 往前清空图片
    }
    for (let i = stop + 2; i < imgList.children.length && setEmptyImg(i); i += 2) { // 往后清空图片
    }
  })

  function init() {
    let n = imgList.children.length > 10 ? 10 : imgList.children.length
    for (let i = 0; i < n; i += 2) {
      const img = imgList.children[i]
      loadImg(img);
    }
  }

  function getShowImgIndex() {
    let scrollTop = document.documentElement.scrollTop; // 页面向上滚动的高度
    let scrollEnd = scrollTop + window.innerHeight; //浏览器自身高度
    let n = 0
    let i = lastIndex
    if (i > unit) {
      i -= unit
    } else {
      i = 0
    }
    for (; i < imgList.children.length; i += 2) {
      const imgTop = imgList.children[i].offsetTop
      const imgEnd = imgList.children[i].offsetTop + imgList.children[i].height
      n++
      if (imgTop < scrollEnd && imgEnd > scrollTop) {
    	document.title = srcTitle + ' 第'+ imgList.children[i].name +'话'
        lastIndex = i
        return i
      }
    }
    return 0
  }

  // 加载图片
  function loadImg(img, emptySrc = false) {
    img.setAttribute('src', img.getAttribute('data-src'))
  }

  // 设置图片为空
  function setEmptyImg(index) {
    const img = imgList.children[index]
    if (img.getAttribute('src')) {
      img.setAttribute('src','')
      return true
    }
    return false
  }
</script>
</html>
`
