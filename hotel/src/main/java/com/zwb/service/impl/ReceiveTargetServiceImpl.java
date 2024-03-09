package com.zwb.service.impl;

import java.util.List;

import com.zwb.mapper.ReceiveTargetMapper;
import com.zwb.page.Page;
import com.zwb.pojo.ReceiveTargetPo;
import com.zwb.service.ReceiveTargetService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;



@Transactional
@Service(value="receiveTargetService")
public class ReceiveTargetServiceImpl implements ReceiveTargetService {

	@Autowired
	private ReceiveTargetMapper receiveTargetMapper;
	
	@Override
	public int deleteById(Integer id) {
		return receiveTargetMapper.deleteById(id);
	}

	@Override
	public int insertAll(ReceiveTargetPo receiveTargetPo) {
		return receiveTargetMapper.insertAll(receiveTargetPo);
	}

	@Override
	public ReceiveTargetPo selectById(Integer id) {
		return receiveTargetMapper.selectById(id);
	}

	@Override
	public int updateById(ReceiveTargetPo receiveTargetPo) {
		return receiveTargetMapper.updateById(receiveTargetPo);
	}

	@Override
	public Page<ReceiveTargetPo> pageFuzzyselect(String teamName, Page<ReceiveTargetPo> vo) {
		int start=0;
		if (vo.getCurrentPage()>1) {
			start=(vo.getCurrentPage()-1)*vo.getPageSize();
		}
		List<ReceiveTargetPo> list= receiveTargetMapper.pageFuzzyselect(teamName, start, vo.getPageSize());
		vo.setResult(list);
		int count= receiveTargetMapper.countFuzzyselect(teamName);
		vo.setTotal(count);
		return vo;
	}

	@Override
	public List<ReceiveTargetPo> ajaxSelect(String teamName) {
		return receiveTargetMapper.ajaxSelect(teamName);
	}

	@Override
	public int selectYZ(String teamCode) {
		return this.receiveTargetMapper.selectYZ(teamCode);
	}

}
