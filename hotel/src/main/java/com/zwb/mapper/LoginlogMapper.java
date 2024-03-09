package com.zwb.mapper;

import com.zwb.pojo.Loginlog;
import com.zwb.pojo.LoginlogExample;
import org.apache.ibatis.annotations.Param;
import org.springframework.stereotype.Repository;

import java.util.List;
@Repository
public interface LoginlogMapper {
    long countByExample(LoginlogExample example);

    int deleteByExample(LoginlogExample example);

    int deleteByPrimaryKey(Integer id);

    int insert(Loginlog record);

    int insertSelective(Loginlog record);

    List<Loginlog> selectByExample(LoginlogExample example);

    Loginlog selectByPrimaryKey(Integer id);

    int updateByExampleSelective(@Param("record") Loginlog record, @Param("example") LoginlogExample example);

    int updateByExample(@Param("record") Loginlog record, @Param("example") LoginlogExample example);

    int updateByPrimaryKeySelective(Loginlog record);

    int updateByPrimaryKey(Loginlog record);

    //分页模糊查询
    public List<Loginlog> pageFuzzyselect(@Param("loginName")String loginName, @Param("start")int start, @Param("pageSize")int pageSize);

    //分页模糊查询总条数
    public int countFuzzyselect(String loginName);
}