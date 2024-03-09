package com.zwb.service.impl;

import java.util.List;

import com.zwb.mapper.PassengerMapper;
import com.zwb.page.Page;
import com.zwb.pojo.PassengerPo;
import com.zwb.service.PassengerService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;



@Transactional
@Service(value="passengerService")
public class PassengerServiceImpl implements PassengerService {

	@Autowired
	private PassengerMapper passengerMapper;
	
	@Override
	public int deleteById(Integer id) {
		return passengerMapper.deleteById(id);
	}

	@Override
	public int insertAll(PassengerPo passengerPo) {
		return passengerMapper.insertAll(passengerPo);
	}

	@Override
	public PassengerPo selectById(Integer id) {
		return passengerMapper.selectById(id);
	}

	@Override
	public int updateById(PassengerPo passengerPo) {
		return passengerMapper.updateById(passengerPo);
	}

	@Override
	public Page<PassengerPo> pageFuzzyselect(String name, Page<PassengerPo> vo) {
		int start=0;
		if (vo.getCurrentPage()>1) {
			start=(vo.getCurrentPage()-1)*vo.getPageSize();
		}
		List<PassengerPo> list= passengerMapper.pageFuzzyselect(name, start, vo.getPageSize());
		vo.setResult(list);
		int count= passengerMapper.countFuzzyselect(name);
		vo.setTotal(count);
		return vo;
	}

	@Override
	public List<PassengerPo> selectAll() {
		return passengerMapper.selectAll();
	}

	@Override
	public List<PassengerPo> selectAjaxList(String name) {
		return passengerMapper.selectAjaxList(name);
	}

	@Override
	public int selectYZ(String papersNumber) {
		return this.passengerMapper.selectYZ(papersNumber);
	}

}
