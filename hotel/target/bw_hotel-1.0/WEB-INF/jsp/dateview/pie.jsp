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
            text: '散客收益金额扇形图',
            textStyle: {
                fontSize: 25,
                fontWeight: 'bolder',
                color: 'rgba(181,35,96,0.76)'
            }
        },
        backgroundColor: 'rgba(208,243,151,0.25)',
        tooltip: {
            trigger: 'item',
            formatter: "{a} <br/>{b}: {c} ({d}%)",
            textStyle: {    //图例文字的样式
                color: '#fff',  //文字颜色
                fontSize: 20    //文字大小
            }
        },
        // 图例
        legend: {
            // 图例选择的模式，控制是否可以通过点击图例改变系列的显示状态。默认开启图例选择，可以设成 false 关闭。
            selectedMode: true,
            /* orient: 'vertical', */
            /*  x : 'right',   //图例显示在右边
             y: 'center',  */  //图例在垂直方向上面显示居中
            bottom: 0,
            itemWidth: 10,  //图例标记的图形宽度
            itemHeight: 10, //图例标记的图形高度
            data: ['${year12}', '${year11}', '${year10}', '${year9}', '${year8}', '${year7}',
                '${year6}', '${year5}', '${year4}', '${year3}', '${year2}', '${year1}'],
            textStyle: {    //图例文字的样式
                color: '#8d1020',  //文字颜色
                fontSize: 18    //文字大小
            }
        },
        toolbox: {
            show: true,
            feature: {
                mark: {show: true},
                saveAsImage: {show: true},
            }
        },
        graphic: [
            {
                type: 'text',
                top: '40%',
                left: 'center',
                style: {
                    text: '${sZong}',
                    fill: '#38993a',
                    fontSize: 25,
                    fontWeight: 'bolder'
                }
            },
            {
                type: 'text',
                top: '53%',
                left: 'center',
                style: {
                    text: '散客总金额',
                    fill: '#A6A8B6',
                    fontSize: 25,
                    fontWeight: 'normal'
                }
            },
        ],
        //圆环的颜色
        color: ['#ff7f50', '#87cefa', '#da70d6', '#32cd32',
            '#6495ed', '#1e90ff', '#ff6347', '#7b68ee',
            '#00fa9a', '#ffd700', 'rgba(112,102,255,0.32)', '#ff6666'],
        series: [
            {
                name: '散客收入占比图',
                type: 'pie',
                center: ['50%', '45%'], //饼图的中心（圆心）坐标，数组的第一项是横坐标，第二项是纵坐标。
                radius: ['35%', '60%'],//饼图的半径，数组的第一项是内半径，第二项是外半径。[ default: [0, '75%'] ]
                avoidLabelOverlap: false,
                /*  animation: false, */  //是否禁用动画效果
                label: {
                    normal: {
                        show: true,//是否显示标签
                        //标签的位置。'outside'饼图扇区外侧，通过视觉引导线连到相应的扇区。'inside','inner' 同 'inside',饼图扇区内部。'center'在饼图中心位置。
                        position: 'left',
                        //显示的标签的内容
                        //a（系列名称），b（数据项名称），c（数值）, d（百分比）
                        formatter: "{a},{b}:{c}（{d}%）",

                    },
                    emphasis: {
                        //鼠标放在圆环上显示的标签样式
                        show: true,
                        textStyle: {
                            fontSize: '20',
                            fontWeight: 'bold'
                        }
                    }
                },
                labelLine: {
                    normal: {
                        show: true,//是否显示引导线
                        length: 10,  //百分比引导线
                        length2: 20    //视觉引导项第二段的长度。
                    }
                },
                itemStyle: {
                    normal: {
                        borderColor: "#FFFFFF",
                        borderWidth: 1,
                        label: {
                            show: true,
                            formatter: '{d}%'
                        },
                    }
                },
                data: [
                    {
                        value:${sZongFeiYong12},
                        name: '${year12}'
                    },
                    {
                        value:${sZongFeiYong11},
                        name: '${year11}'
                    },
                    {
                        value:${sZongFeiYong10},
                        name: '${year10}'
                    },
                    {
                        value:${sZongFeiYong9},
                        name: '${year9}'
                    }, {
                        value:${sZongFeiYong8},
                        name: '${year8}'
                    }, {
                        value:${sZongFeiYong7},
                        name: '${year7}'
                    }, {
                        value:${sZongFeiYong6},
                        name: '${year6}'
                    }, {
                        value:${sZongFeiYong5},
                        name: '${year5}'
                    }, {
                        value:${sZongFeiYong4},
                        name: '${year4}'
                    },
                    {
                        value:${sZongFeiYong3},
                        name: '${year3}'
                    },
                    {
                        value:${sZongFeiYong2},
                        name: '${year2}'
                    },
                    {
                        value:${sZongFeiYong1},
                        name: '${year1}'
                    }
                ]

            }
        ]
    };

    // 使用刚指定的配置项和数据显示图表。
    myChart.setOption(option);
</script>
</body>
</html>
