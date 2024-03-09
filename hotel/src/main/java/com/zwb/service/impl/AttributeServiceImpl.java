package com.zwb.service.impl;

import java.util.List;

import com.zwb.mapper.AttributeMapper;
import com.zwb.pojo.AttributePo;
import com.zwb.service.AttributeService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;



@Transactional
@Service(value="attributeService")
public class AttributeServiceImpl implements AttributeService {

	@Autowired
	private AttributeMapper attributeMapper;
	
	@Override
	public List<AttributePo> selectGuestRoomLevel() {
		return attributeMapper.selectGuestRoomLevel();
	}

	@Override
	public List<AttributePo> selectRoomState() {
		return attributeMapper.selectRoomState();
	}

	@Override
	public List<AttributePo> selectCommodityType() {
		return attributeMapper.selectCommodityType();
	}

	@Override
	public List<AttributePo> selectUOM() {
		return attributeMapper.selectUOM();
	}

	

	@Override
	public int deleteById(Integer id) {
		return attributeMapper.deleteById(id);
	}

	@Override
	public int insertAll(int newid, String newname) {
		return attributeMapper.insertAll(newid, newname);
	}

	@Override
	public List<AttributePo> selectTargetType() {
		return attributeMapper.selectTargetType();
	}

	@Override
	public List<AttributePo> selectGender() {
		return attributeMapper.selectGender();
	}

	@Override
	public List<AttributePo> selectNation() {
		return attributeMapper.selectNation();
	}

	@Override
	public List<AttributePo> selectPassengerLevel() {
		return attributeMapper.selectPassengerLevel();
	}

	@Override
	public List<AttributePo> selectEducationDegree() {
		return attributeMapper.selectEducationDegree();
	}

	@Override
	public List<AttributePo> selectPapers() {
		return attributeMapper.selectPapers();
	}

	@Override
	public List<AttributePo> selectThingReason() {
		return attributeMapper.selectThingReason();
	}

	@Override
	public List<AttributePo> selectPassengerType() {
		return attributeMapper.selectPassengerType();
	}

	@Override
	public List<AttributePo> selectBillUnit() {
		return attributeMapper.selectBillUnit();
	}

	@Override
	public List<AttributePo> selectPayWay() {
		return attributeMapper.selectPayWay();
	}

	@Override
	public List<AttributePo> selectRentOutType() {
		return attributeMapper.selectRentOutType();
	}

	@Override
	public List<AttributePo> selectIsPay() {
		return attributeMapper.selectIsPay();
	}

	@Override
	public List<AttributePo> selectPredetermineState() {
		return attributeMapper.selectPredetermineState();
	}

}
