<%@ page language="java" import="java.util.*" pageEncoding="UTF-8" %>
<%@ taglib uri="http://java.sun.com/jsp/jstl/core" prefix="c" %>
<%@ taglib prefix="fmt" uri="http://java.sun.com/jsp/jstl/fmt" %>

<c:set value="${pageContext.request.contextPath}" scope="page" var="ctx"></c:set>
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
<head>

    <title></title>

    <meta http-equiv="pragma" content="no-cache">
    <meta http-equiv="cache-control" content="no-cache">
    <meta http-equiv="expires" content="0">
    <meta http-equiv="keywords" content="keyword1,keyword2,keyword3">
    <meta http-equiv="description" content="This is my page">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">

    <!--
    <link rel="stylesheet" type="text/css" href="styles.css">
    -->
    <script type="text/javascript" src="${ctx}/js/echarts/echarts.min.js"></script>
</head>
<style>

</style>
<body>

<div id="main" style="width: 98%;height:82%;"></div>


<script type="text/javascript">
    // 基于准备好的dom，初始化echarts实例
    var myChart = echarts.init(document.getElementById('main'));

    // 指定图表的配置项和数据
    var option = {
        title: {
            text: '收益金额折线图',
            textStyle:{
                fontSize :25,
                fontWeight:'bolder',
                color:'rgba(45,151,200,0.86)'
            }
        },
        backgroundColor: 'rgba(151,243,243,0.25)',
        tooltip: {
            trigger: 'axis'
        },
        legend: {
            data: ['散客', '团队'],
            textStyle: {
                fontSize: 20
            }
        },
        toolbox: {
            feature: {
                saveAsImage: {}
            }
        },
        grid: {
            left: '3%',
            right: '4%',
            bottom: '3%',
            containLabel: true
        },
        xAxis: {
            type: 'category',
            boundaryGap: false,
            data: ['${year12}', '${year11}', '${year10}', '${year9}', '${year8}', '${year7}',
                '${year6}', '${year5}', '${year4}', '${year3}', '${year2}', '${year1}'],
            axisLabel: {
                show: true,
                textStyle: {
                    fontSize: 15,
                    color: 'rgba(10,10,10,0.95)'
                }
            }
        },
        yAxis: [{
            type: 'value'
        }],
        series: [
            {
                name: '散客',
                type: 'line',
                stack: '总量',
                label: {
                    normal: {
                        show: true,
                        textStyle: {
                            color: 'rgba(252,5,75,0.66)',
                            fontSize: 18,
                            fontWeight: 'bold'
                        }
                    }
                },
                itemStyle: {
                    normal: {
                        areaStyle: {
                            color: 'rgba(3,246,101,0.55)'
                        },
                        lineStyle: {
                            color: 'rgba(25,49,59,0.74)'
                        }
                    }
                },
                data: [${sZongFeiYong12}, ${sZongFeiYong11}, ${sZongFeiYong10}, ${sZongFeiYong9},
                    ${sZongFeiYong8}, ${sZongFeiYong7}, ${sZongFeiYong6}, ${sZongFeiYong5},
                    ${sZongFeiYong4}, ${sZongFeiYong3}, ${sZongFeiYong2}, ${sZongFeiYong1}]
            },
            {
                name: '团队',
                type: 'bar',
                stack: '总量',
                barWidth: 55,//柱形图宽度
                label: {
                    normal: {
                        show: true,
                        position: 'top',
                        textStyle: {
                            color: '#0e0e0e',
                            fontSize: 18,
                            fontWeight: 'bold'
                        }
                    }
                },
                itemStyle: {
                    normal: {
                        areaStyle: {
                            color: 'rgba(75,96,96,0.52)'
                        },
                        lineStyle: {
                            color: 'rgba(64,200,33,0.66)'
                        }
                    }
                },
                data: [${tZongFeiYong12}, ${tZongFeiYong11}, ${tZongFeiYong10}, ${tZongFeiYong9},
                    ${tZongFeiYong8}, ${tZongFeiYong7}, ${tZongFeiYong6}, ${tZongFeiYong5},
                    ${tZongFeiYong4}, ${tZongFeiYong3}, ${tZongFeiYong2}, ${tZongFeiYong1}]
            }
        ]
    };

    // 使用刚指定的配置项和数据显示图表。
    myChart.setOption(option);
</script>
</body>
</html>
