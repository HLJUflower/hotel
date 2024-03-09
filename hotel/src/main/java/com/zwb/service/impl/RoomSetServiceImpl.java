package com.zwb.service.impl;

import java.util.List;

import com.zwb.mapper.RoomSetMapper;
import com.zwb.page.Page;
import com.zwb.pojo.RoomSetPo;
import com.zwb.service.RoomSetService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;



@Transactional
@Service(value="roomSetService")
public class RoomSetServiceImpl implements RoomSetService {

	@Autowired
	private RoomSetMapper roomSetMapper;
	
	/*@Override
	public List<RoomSetPo> selectAll() {
		return roomSetDao.selectAll();
	}
*/
	@Override
	public int deleteById(Integer id) {
		return roomSetMapper.deleteById(id);
	}

	@Override
	public int insertAll(RoomSetPo roomSetPo) {
		return roomSetMapper.insertAll(roomSetPo);
	}

	@Override
	public RoomSetPo selectById(Integer id) {
		return roomSetMapper.selectById(id);
	}

	@Override
	public int updateById(RoomSetPo roomSetPo) {
		return roomSetMapper.updateById(roomSetPo);
	}

	//模糊查询
	/*@Override
	public List<RoomSetPo> fuzzyselectPo(String roomNumber) {
		return roomSetDao.fuzzyselectPo(roomNumber);
	}*/
   

//分页的模糊查询
	@Override
	public Page<RoomSetPo> pageFuzzyselect(String roomNumber, Page<RoomSetPo> vo) {
		int start=0;
		if (vo.getCurrentPage()>1) {
			start=(vo.getCurrentPage()-1)*vo.getPageSize();
		}
		List<RoomSetPo> list=this.roomSetMapper.pageFuzzyselect(roomNumber, start, vo.getPageSize());
		vo.setResult(list);
		
		int count=this.roomSetMapper.countFuzzyselect(roomNumber);
		vo.setTotal(count);
		return vo;
	}

	@Override
	public List<RoomSetPo> selectAll() {
		return roomSetMapper.selectAll();
	}

	@Override
	public List<RoomSetPo> selectByLeveId(Integer id) {
		return roomSetMapper.selectByLeveId(id);
	}

	@Override
	public int updateByIdToRoomState(RoomSetPo roomSetPo) {
		return roomSetMapper.updateByIdToRoomState(roomSetPo);
	}

	@Override
	public List<RoomSetPo> selectInformation(String roomNumber) {
		return roomSetMapper.selectInformation(roomNumber);
	}

	@Override
	public List<RoomSetPo> levelSelectInformation(Integer guestRoomLevelID) {
		return this.roomSetMapper.levelSelectInformation(guestRoomLevelID);
	}

	@Override
	public int selectYZ(String roomNumber) {
		return this.roomSetMapper.selectYZ(roomNumber);
	}

	@Override
	public int selectLevelById(Integer id) {
		return roomSetMapper.selectLevelById(id);
	}

	@Override
	public int countRoomByLevel(Integer guestRoomLevelID) {
		return roomSetMapper.countRoomByLevel(guestRoomLevelID);
	}


}
