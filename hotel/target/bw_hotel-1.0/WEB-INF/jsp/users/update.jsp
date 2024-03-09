<%@ page language="java" import="java.util.*" pageEncoding="UTF-8"%>
<%@ taglib uri="http://java.sun.com/jsp/jstl/core" prefix="c" %>
<%@ taglib prefix="fmt" uri="http://java.sun.com/jsp/jstl/fmt" %>

<c:set  value="${pageContext.request.contextPath}" scope="page" var="ctx"></c:set>
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
  <head>
    
    <title></title>
    
	<meta http-equiv="pragma" content="no-cache">
	<meta http-equiv="cache-control" content="no-cache">
	<meta http-equiv="expires" content="0">    
	<meta http-equiv="keywords" content="keyword1,keyword2,keyword3">
	<meta http-equiv="description" content="This is my page">
	<meta name="viewport" content="width=device-width, initial-scale=1.0" />

	
<!--   <link rel="stylesheet" href="${ctx}/css/roomset/roomset.css" type="text/css"></link> -->
  <link rel="stylesheet" href="${ctx}/bootstrap/css/bootstrap.css" type="text/css"></link>
  <link href="${ctx}/bootstrap/css/bootstrap-responsive.css" rel="stylesheet">  <!-- start 响应式布局要添加的 -->
  <script src="${ctx}/bootstrap/js/jquery-3.1.1.min.js"></script>
  <script src="${ctx}/bootstrap/js/bootstrap.js"></script>
	  <link href="${ctx}/css/backitem.css" rel="stylesheet" type="text/css">
   
   <style>
   
   .container{
     margin-top: 10px;
   }
   
   .theadone{
     background-color: #438eb9; color: #fff;
    }
   
   
   .labelroomnumber{
      position: relative;
      font-size: 17px;
      float: left;
      margin-top: 15px;
    }
    
    .marginrightone{
      margin-right: 33px;
    }
    
    .marginrighttwo{
      margin-right: 16.5px;
    }
    
    .marginrightthree{
      margin-right: 43px;
    }
    
   .textone{
    margin-top:12px;
    }
    
    .inputone{
    width:54.2%;
    }
    
    .inputtwo{
      width:46.8%;
    }
    
    .inputthree{
      width:46.8%;
    }
    
    .cboone{
      margin-top:10px;
      height: 27px;
    }
    
    .margin-top-one{
      margin-top: -10px;
    }
    
    .margin-top-two{
      margin-top: -20px;
    }
  
    .textwidth{
      width: 138px;
    }
    
    .radiusone{
     margin-left: -4px;
    }
    
     @media(min-width:731px) and (max-width:767px){
      .inputthree{
      width: 50.5%;
    }
    }
    
    @media(min-width:550px) and (max-width:730px){
      .inputtwo{
      width: 46.2%;
    }
        .inputthree{
      width: 49.2%;
    }
    }
    
    
     @media(min-width:431px) and (max-width:550px){
      .inputtwo{
      width: 43.2%;
    }
      .inputthree{
      width: 47.5%;
    }
    }
    
    @media(min-width:366px) and (max-width:430px){
      .inputtwo{
      width: 40.2%;
    }
      .inputthree{
      width: 46%;
    }
    }
    
     @media(min-width:285px) and (max-width:366px){
      .inputtwo{
      width: 37.2%;
    }
      .inputthree{
      width: 44%;
    }
    }
    
     @media(min-width:237px) and (max-width:285px){
      .inputtwo{
      width: 30%;
    }
      .inputthree{
      width: 40%;
    }
    }
    .span12{
      width:95%;
    }
    .yansered{
      color:red;
    }
  
  </style>
  
  </head>
  
 
  <body>
  <div class="container" style="height:630px;overflow-x:auto;">
  
    <div class="span11" style=" border: solid; border-color: #DDDDDD;">
    <div class="span9 margin-top-one">
      <div class="row-fluid">
        <h3 style="text-align: center;">管理员修改</h3>
        
      </div>
    </div>
    
    <form action="${ctx}/users/update.do" method="post" onsubmit="return verify()">
    <!--  ———————————————————————————————————————————————————————————————————————————————————————— -->
		<input type="hidden" name="id" value="${user.id}" />
	    <div class="span12">
	      <div class="row-fluid">
		     <div class="span3">
		        用户名：
		        <input id="nameId" name="username" type="text" value="${user.username}" style="width:97%;height:27px;float:left;" onchange="onchangeOne()" onblur="YZ(this.value)">
		        <div id="divOne" style="float:right;">
			         <label class="yansered" style="margin-top:7px;">*</label>
			    </div> 
		      </div>
		  </div>
	    </div>
	    <!--  ———————————————————————————————————————————————————————————————————————————————————————— -->
	    <div class="span12">
	      <div class="row-fluid">
		     <div class="span3">
		        密码：
		        <input id="pwd" name="password" type="password" value="${user.password}" style="width:97%;height:27px;float: left" onchange="onchangeOne()">
				 <div id="divTwo" style="float:right;">
					 <label class="yansered" style="margin-top:7px;">*</label>
				 </div>
		     </div>
		  </div>
			<label class="yansered" style="margin-top:7px;">密码要求：至少8个字符,至少1个大写字母，1个小写字母和1个数字,不能包含特殊字符（非数字字母）</label>
	    </div>
	    <!--  ———————————————————————————————————————————————————————————————————————————————————————— -->
		<div class="span12">
			<div class="row-fluid">
				<div class="span3">
					确认密码：<input id="repwd" name="repassword" type="password" value="${user.password}" style="width:97%;height:27px;float: left" onchange="onchangeOneRe()" />
					<div id="divTwoo" style="float:right;">
						<label class="yansered" style="margin-top:7px;">*</label>
					</div>
				</div>
			</div>
		</div>
	    <!--  ———————————————————————————————————————————————————————————————————————————————————————— -->
		  <div class="span6" style="text-align:center;">
		      <div class="row-fluid">
			      <div class="span12" style="margin-top: 10px;margin-bottom: 8px;">
				   	 <button class="btn btn-primary" type="submit">
				   	 <li class="icon-check icon-white"></li>保存</button>
				  </div> 
				 
			  </div>
	      </div>
      </form>
      
      
      <div class="span4" style="text-align:center;">
      <div class="row-fluid">
		   <div class="span4"  style="margin-top: 10px;margin-bottom: 8px;"> 
		   	 <button class="btn btn-warning" type="button" onclick="deletefunction()">
		   	  <li class="icon-remove icon-white"></li>取消</button>
		   </div>
	     </div>
      </div>
       
    </div>
  </div>
  </body>
 
 <script type="text/javascript">


    function verify(){
	    if(document.getElementById("nameId").value==""){
	       alert("用户名  是必填项，不能为空哦！");
	       document.getElementById("nameId").focus();
	       return false;
	    }else if(document.getElementById("pwd").value==""){
	       alert("密码  是必填项，不能为空哦！");
	       document.getElementById("pwd").focus();
	       return false;
	    }else if(document.getElementById("repwd").value==""){
	     alert("确认密码  是必填项，不能为空哦！");
	     document.getElementById("repwd").focus();
	     return false;
        }else{
	       return true;
	    }
   }
   
    function deletefunction(){
     parent.document.getElementById('Mainid').src='${ctx}/users/tolist.do';
   }
   //加载以后即判断
   $(function () {
	   if(document.getElementById("nameId").value!=""){
		   document.getElementById("divOne").style.display="none";
	   }else{
		   document.getElementById("divOne").style.display="block"; //显示
	   }
	   //密码
	   if(document.getElementById("pwd").value!=""){
		   document.getElementById("divTwo").style.display="none";
	   }else{
		   document.getElementById("divTwo").style.display="block"; //显示
	   }
	   if(document.getElementById("repwd").value!=""){
		   document.getElementById("divTwoo").style.display="none";
	   }else{
		   document.getElementById("divTwoo").style.display="block"; //显示
	   }

   })
   function onchangeOne(){
     //用户名
     if(document.getElementById("nameId").value!=""){
       document.getElementById("divOne").style.display="none";
     }else{
       document.getElementById("divOne").style.display="block"; //显示
     }

     //密码
     if(document.getElementById("pwd").value!=""){
		 document.getElementById("divTwo").style.display="none";
     	//至少8个字符，至少1个大写字母，1个小写字母和1个数字,不能包含特殊字符（非数字字母）
        if(!document.getElementById("pwd").value.trim().match("^(?=.*[A-Za-z])(?=.*\\d)[A-Za-z\\d]{8,}$"))
	       {
	          document.getElementById("pwd").focus();
	          document.getElementById("pwd").value="";
	          alert("密码不合法！");
	       }
     }else{
		 document.getElementById("divTwo").style.display="block"; //显示
	 }
   }
   function onchangeOneRe() {
			//确认密码
	   if(document.getElementById("repwd").value!=""){
		   document.getElementById("divTwoo").style.display="none";
		   if(document.getElementById("pwd").value!=document.getElementById("repwd").value){
			   alert("两次密码不一致");
		   }
	   }else{
		   document.getElementById("divTwoo").style.display="block"; //显示
	   }
   }
   
   
   function YZ(value){
   console.log(11);
     if(value!=""){
       $.ajax({                                                      
          cache:false,                                             //是否使用缓存提交 如果为TRUE 会调用浏览器的缓存 而不会提交
          type: "POST",                                           //上面3行都是必须要的
          url: '${ctx}/users/YZ.do',       //地址 type 带参数
          data:"userName="+value,                         // IDCardValue 自定义的。相当于name把值赋予给 他可以在servlet 获取
          async:false,                                          // 是否 异步 提交
          success: function (result) {                          // 不出现异常 进行立面方
              if(result>=1){                                  
                   alert("此用户名已经存在！");                     //提示框
                   document.getElementById("nameId").value="";     //这个id的文本框的值 将会清空
                   document.getElementById("nameId").focus();      // 给这个id的文本框提供焦点
                   document.getElementById("divOne").style.display="block"; //显示
              }
          },
          error: function(data) {  }
      });
     }
   }
 </script>
</html>
