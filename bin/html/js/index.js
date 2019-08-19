
var iMaxPageNum = 0;
function getquerystring(name)
{
	var reg = new RegExp("(^|&)"+ name +"=([^&]*)(&|$)");
	var r = window.location.search.substr(1).match(reg);
	if(r!=null)return  unescape(r[2]); return null;
}

function setmaxpagenum()
{
	iMaxPageNum = document.getElementById("maxPage").innerText;
}

function prepage()
{
	var iCurPage = getquerystring("page");
	if(iCurPage == null){
		iCurPage = 1;
	}else{
		iCurPage--;
	}
	setcurpage(iCurPage);
}

function nextpage()
{
	var iCurPage = getquerystring("page");
	if(iCurPage == null){
		iCurPage = 1;
	}else{
		iCurPage++;
	}
	setcurpage(iCurPage);
}


function setcurpage(idx)
{
	if(idx < 1){
		idx = 1;
	}else if(idx > iMaxPageNum){
		idx = iMaxPageNum;
	}
	var pageAddr = "?page=" + idx;
	window.location.href = pageAddr;
}

function initpage()
{
	setmaxpagenum();
	convert2md();
}

function convert2md()
{
	//var list = document.getElementsByTagName("li");
	var list = document.getElementsByName("markdownContent");
	for(var i=0; i < list.length; i++){
		//alert(list[i].innerHTML);
		compile(list[i]);
	}
}

function compile(elem){
	//alert(elem.innerHTML);
	var converter = new showdown.Converter();
	//alert(converter);
	var html = converter.makeHtml(elem.textContent);
	//alert(html);
	elem.innerHTML = html;
}
