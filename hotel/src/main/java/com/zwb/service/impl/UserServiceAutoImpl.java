package com.zwb.service.impl;

import com.zwb.mapper.UsersMapperAuto;
import com.zwb.page.Page;
import com.zwb.pojo.Users;
import com.zwb.pojo.UsersExample;
import com.zwb.service.UserServieAuto;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import javax.annotation.Resource;
import java.util.List;

/**
 * @Description:
 */
@Transactional
@Service(value = "userServiceAuto")
public class UserServiceAutoImpl implements UserServieAuto {
    @Resource
    private UsersMapperAuto usersMapperAuto;


    @Override
    public int selectIsExit(String userName) {
        UsersExample usersExample = new UsersExample();
        usersExample.createCriteria().andUsernameEqualTo(userName);
        List<Users> users = usersMapperAuto.selectByExample(usersExample);
        return users.size();
    }

    @Override
    public int deleteByUserId(Integer id) {
        int i = usersMapperAuto.deleteByPrimaryKey(id);
        return i;
    }

    @Override
    public int insertUsers(Users users) {
        int insert = usersMapperAuto.insert(users);

        return insert;
    }

    @Override
    public int updateByUserName(Users user) {
        return usersMapperAuto.updateByPrimaryKey(user);
    }

    @Override
    public Users selectUserById(Integer id) {
        return usersMapperAuto.selectByPrimaryKey(id);
    }

    @Override
    public List<Users> selectByExample(UsersExample example) {
        return usersMapperAuto.selectByExample(example);
    }

    @Override
    public Page<Users> pageFuzzyselect(String name, Page<Users> vo) {
        int start =0;
        if(vo.getCurrentPage()>1){
            start=(vo.getCurrentPage()-1)*vo.getPageSize();
        }
        List<Users> usersList = usersMapperAuto.pageFuzzyselect(name, start, vo.getPageSize());
        vo.setResult(usersList);
        int count = usersMapperAuto.countFuzzyselect(name);
        vo.setTotal(count);
        return vo;
    }

    @Override
    public List<Users> selectAll() {
        return usersMapperAuto.selectByExample(null);
    }

    @Override
    public List<Users> selectAjaxList(String userName) {
        UsersExample usersExample = new UsersExample();
        usersExample.createCriteria().andUsernameLike(userName);


        return usersMapperAuto.selectByExample(usersExample);
    }
}
