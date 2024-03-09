package com.zwb.mapper;

import com.zwb.pojo.Users;
import com.zwb.pojo.UsersExample;
import org.apache.ibatis.annotations.Param;
import org.springframework.stereotype.Repository;

import java.util.List;
@Repository
public interface UsersMapperAuto {
    long countByExample(UsersExample example);

    int deleteByExample(UsersExample example);

    int deleteByPrimaryKey(Integer id);

    int insert(Users record);

    int insertSelective(Users record);

    List<Users> selectByExample(UsersExample example);

    Users selectByPrimaryKey(Integer id);

    int updateByExampleSelective(@Param("record") Users record, @Param("example") UsersExample example);

    int updateByExample(@Param("record") Users record, @Param("example") UsersExample example);

    int updateByPrimaryKeySelective(Users record);

    int updateByPrimaryKey(Users record);

    //分页模糊查询
    public List<Users> pageFuzzyselect(@Param("name")String name,
                                             @Param("start")int start, @Param("pageSize")int pageSize);

    //分页模糊查询总条数
    public int countFuzzyselect(@Param("name")String name);
}