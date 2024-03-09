package com.zwb.service.impl;

import com.zwb.mapper.LoginlogMapper;
import com.zwb.page.Page;
import com.zwb.pojo.Loginlog;
import com.zwb.service.LoginLogService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

/**
 * Created with IntelliJ IDEA.
 *
 * @Auther: wzy
 * @Date: 2021/04/05/19:52
 * @Description:
 */
@Service
public class LoginLogServiceImpl implements LoginLogService {
    @Autowired
    private LoginlogMapper loginlogMapper;
    @Override
    public int saveLoginLog(Loginlog loginlog) {
       return loginlogMapper.insertSelective(loginlog);
    }

    @Override
    public List<Loginlog> selectAll() {
        return loginlogMapper.selectByExample(null);
    }

    @Override
    public Page<Loginlog> pageFuzzyselect(String loginName, Page<Loginlog> vo) {
        int start=0;
        if (vo.getCurrentPage()>1) {
            start=(vo.getCurrentPage()-1)*vo.getPageSize();
        }
        List<Loginlog> list=loginlogMapper.pageFuzzyselect(loginName, start, vo.getPageSize());
        vo.setResult(list);

        int count=this.loginlogMapper.countFuzzyselect(loginName);
        vo.setTotal(count);
        return vo;
    }

    @Override
    public void updateLoginLogById(Loginlog loginlog) {
        loginlogMapper.updateByPrimaryKey(loginlog);
    }

    @Override
    public Loginlog selectOne(Integer id) {
        return loginlogMapper.selectByPrimaryKey(id);
    }

    @Override
    public void deleteByUserId(int parseInt) {
        loginlogMapper.deleteByPrimaryKey(parseInt);
    }

}
