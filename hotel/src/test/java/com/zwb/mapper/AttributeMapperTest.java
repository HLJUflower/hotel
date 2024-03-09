package com.zwb.mapper;

import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.test.context.ContextConfiguration;
import org.springframework.test.context.junit4.SpringRunner;

@RunWith(SpringRunner.class)
@ContextConfiguration({"classpath:applicationContext-mybatis.xml",
        "classpath:applicationContext-spring.xml",
        "classpath:applicationContext-trans.xml",
        "classpath:springmvc.xml"})
public class AttributeMapperTest {

    @Autowired
    AttributeMapper attributeMapper;

    @Test
    public void test(){

        attributeMapper.insertAll(18,"测试");

    }

}
