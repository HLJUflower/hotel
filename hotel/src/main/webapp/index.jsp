<%@ page language="java" import="java.util.*" pageEncoding="UTF-8"%>
<%
String path = request.getContextPath();
String basePath = request.getScheme()+"://"+request.getServerName()+":"+request.getServerPort()+path+"/";
%>

<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
  <head>
    <base href="<%=basePath%>">
    
    <title>My JSP 'index.jsp' starting page</title>
	<meta http-equiv="pragma" content="no-cache">
	<meta http-equiv="cache-control" content="no-cache">
	<meta http-equiv="expires" content="0">    
	<meta http-equiv="keywords" content="keyword1,keyword2,keyword3">
	<meta http-equiv="description" content="This is my page">
	<meta http-equiv="refresh" content="0;url=<%= path %>/StayRegister/roomList.do">
	<!--
	<link rel="stylesheet" type="text/css" href="styles.css">
	-->
      <!--   <link rel="stylesheet" href="${ctx}/css/roomset/roomset.css" type="text/css"></link> -->
      <link rel="stylesheet" href="${ctx}/bootstrap/css/bootstrap.css" type="text/css"></link>
      <link href="${ctx}/bootstrap/css/bootstrap-responsive.css" rel="stylesheet">  <!-- start 响应式布局要添加的 -->
      <script src="${ctx}/bootstrap/js/jquery-3.1.1.min.js"></script>
      <script src="${ctx}/bootstrap/js/bootstrap.js"></script>
      <link rel="stylesheet" href="${ctx}/bootstrap/css/bootstrap.min.css" type="text/css"></link>

      <script type="text/javascript" src="${ctx}/My97DatePicker/WdatePicker.js"></script>
  </head>
</html>
