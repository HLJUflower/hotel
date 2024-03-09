package com.zwb.service.impl;

import java.sql.Timestamp;
import java.util.List;

import com.zwb.mapper.StayRegisterMapper;
import com.zwb.page.Page;
import com.zwb.pojo.StayRegisterPo;
import com.zwb.service.StayRegisterService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;


@Transactional
@Service(value="stayRegisterService")
public class StayRegisterServiceImpl implements StayRegisterService {

	@Autowired
	private StayRegisterMapper stayRegisterMapper;
	
	@Override
	public int insertAll(StayRegisterPo stayRegisterPo) {
		return stayRegisterMapper.insertAll(stayRegisterPo);
	}

	@Override
	public List<Integer> selectAllRoomNum(Timestamp payTime) {
		return stayRegisterMapper.selectAllRoomNum(payTime);
	}

	@Override
	public Page<StayRegisterPo> pageFuzzyselectOne(int receiveTargeTypeID,
												   int isBillID, String roomNumber, Page<StayRegisterPo> vo) {
		int start=0;
		if (vo.getCurrentPage()>1) {
			start=(vo.getCurrentPage()-1)*vo.getPageSize();
		}
		List<StayRegisterPo> list= stayRegisterMapper.pageFuzzyselectOne(receiveTargeTypeID, isBillID, roomNumber,start, vo.getPageSize());
		vo.setResult(list);
		int count= stayRegisterMapper.countFuzzyselectOne(receiveTargeTypeID, isBillID,roomNumber);
		vo.setTotal(count);
		return vo;
	}

	@Override
	public StayRegisterPo selectById(Integer id) {
		return this.stayRegisterMapper.selectById(id);
	}

	@Override
	public int insertStayregisterdetails(int stayRegisterID, int passengerID) {
		return stayRegisterMapper.insertStayregisterdetails(stayRegisterID, passengerID);
	}

	@Override
	public int insertDeposit(StayRegisterPo stayRegisterPo) {
		return stayRegisterMapper.insertDeposit(stayRegisterPo);
	}

	@Override
	public List<StayRegisterPo> selectDepositById(int id) {
		return stayRegisterMapper.selectDepositById(id);
	}

	@Override
	public int insertConsumptiondetails(StayRegisterPo stayRegisterPo) {
		return stayRegisterMapper.insertConsumptiondetails(stayRegisterPo);
	}

	@Override
	public Page<StayRegisterPo> pageConsumption(int consumptionStayRegisterID,Page<StayRegisterPo> vo) {
		int start=0;
		if (vo.getCurrentPage()>1) {
			start=(vo.getCurrentPage()-1)*vo.getPageSize();
		}
		List<StayRegisterPo> list= stayRegisterMapper.pageConsumption(consumptionStayRegisterID, start, vo.getPageSize());
		vo.setResult(list);
		int count= stayRegisterMapper.countConsumption(consumptionStayRegisterID);
		vo.setTotal(count);
		return vo;
	}

	@Override
	public int deleteConsumption(Integer id) {
		return stayRegisterMapper.deleteConsumption(id);
	}

	@Override
	public int updateSumConst(int id, Double sumConst) {
		return stayRegisterMapper.updateSumConst(id, sumConst);
	}

	@Override
	public List<StayRegisterPo> selectMoney(int id) {
		return stayRegisterMapper.selectMoney(id);
	}

	@Override
	public List<StayRegisterPo> selectChangRoom(int id) {
		return stayRegisterMapper.selectChangRoom(id);
	}

	@Override
	public List<StayRegisterPo> selectAll() {
		return stayRegisterMapper.selectAll();
	}

	@Override
	public int updateRemind(int id, int remind) {
		return stayRegisterMapper.updateRemind(id, remind);
	}

	@Override
	public int updateChangRoom(StayRegisterPo stayRegisterPo) {
		return stayRegisterMapper.updateChangRoom(stayRegisterPo);
	}

	@Override
	public int pay(int id, String remarks, Timestamp payTime, int payWay) {
		return stayRegisterMapper.pay(id, remarks, payTime, payWay);
	}

	@Override
	public Page<StayRegisterPo> pageFuzzyselectTwo(int receiveTargetID,
			int isBillID, String roomNumber, Page<StayRegisterPo> vo) {
		int start=0;
		if (vo.getCurrentPage()>1) {
			start=(vo.getCurrentPage()-1)*vo.getPageSize();
		}
		List<StayRegisterPo> list= stayRegisterMapper.pageFuzzyselectTwo(receiveTargetID, isBillID, roomNumber,start, vo.getPageSize());
		vo.setResult(list);
		int count= stayRegisterMapper.countFuzzyselectTwo(receiveTargetID, isBillID, roomNumber);
		vo.setTotal(count);
		return vo;
	}

	@Override
	public List<StayRegisterPo> selectFormTeamId(int receiveTargetID,
			int isBillID) {
		return this.stayRegisterMapper.selectFormTeamId(receiveTargetID, isBillID);
	}

	@Override
	public List<StayRegisterPo> selectTeamDeposit(int receiveTargetID) {
		return this.stayRegisterMapper.selectTeamDeposit(receiveTargetID);
	}

	@Override
	public List<StayRegisterPo> selectTeamConsumption(int receiveTargetID) {
		return this.stayRegisterMapper.selectTeamConsumption(receiveTargetID);
	}

	@Override
	public List<StayRegisterPo> ajaxSelectTeamRoom(int receiveTargetID,
			String roomNumber) {
		return this.stayRegisterMapper.ajaxSelectTeamRoom(receiveTargetID, roomNumber);
	}

	@Override
	public List<StayRegisterPo> ajaxSelectTeamFormTime(int receiveTargetID,
			Timestamp min, Timestamp max) {
		return this.stayRegisterMapper.ajaxSelectTeamFormTime(receiveTargetID, min, max);
	}

	@Override
	public List<StayRegisterPo> ajaxSelectTeamDeposit(int receiveTargetID,
			Timestamp min, Timestamp max) {
		return this.stayRegisterMapper.ajaxSelectTeamDeposit(receiveTargetID, min, max);
	}

	@Override
	public List<StayRegisterPo> ajaxSelectTeamConsumption(int receiveTargetID,
			Timestamp min, Timestamp max) {
		return this.stayRegisterMapper.ajaxSelectTeamConsumption(receiveTargetID, min, max);
	}

	@Override
	public List<StayRegisterPo> selectDepositJinJianBan(int id) {
		return this.stayRegisterMapper.selectDepositJinJianBan(id);
	}

	@Override
	public List<StayRegisterPo> selectConsumptionJinJianBan(int id) {
		return this.stayRegisterMapper.selectConsumptionJinJianBan(id);
	}

	@Override
	public StayRegisterPo selectInformationXiangQingBan(int id) {
		return this.stayRegisterMapper.selectInformationXiangQingBan(id);
	}

	@Override
	public int changOverTeam(int id, int receiveTargetID) {
		return this.stayRegisterMapper.changOverTeam(id, receiveTargetID);
	}

	@Override
	public List<StayRegisterPo> selectFormTeamIdTwo(int isBillID) {
		return this.stayRegisterMapper.selectFormTeamIdTwo(isBillID);
	}

	@Override
	public Page<StayRegisterPo> pageFuzzyselectThree(int isBillID,
			String roomNumber, Page<StayRegisterPo> vo) {
		int start=0;
		if (vo.getCurrentPage()>1) {
			start=(vo.getCurrentPage()-1)*vo.getPageSize();
		}
		List<StayRegisterPo> list= stayRegisterMapper.pageFuzzyselectThree(isBillID, roomNumber,start, vo.getPageSize());
		vo.setResult(list);
		int count= stayRegisterMapper.countFuzzyselectThree(isBillID, roomNumber);
		vo.setTotal(count);
		return vo;
	}


	@Override
	public List<StayRegisterPo> selectShuJuTongJi(Timestamp min, Timestamp max) {
		return this.stayRegisterMapper.selectShuJuTongJi(min, max);
	}
	
	
	
	
	
	
	
	
	
	
/*-------------------------------------FinancialStatistics--------------------------------------------------------*/

	@Override
	public Page<StayRegisterPo> pageFuzzyselectFour(Page<StayRegisterPo> vo) {
		int start=0;
		if (vo.getCurrentPage()>1) {
			start=(vo.getCurrentPage()-1)*vo.getPageSize();
		}
		List<StayRegisterPo> list= stayRegisterMapper.pageFuzzyselectFour(start, vo.getPageSize());
		vo.setResult(list);
		int count= stayRegisterMapper.countFuzzyselectFour();
		vo.setTotal(count);
		return vo;
	}

	@Override
	public Page<StayRegisterPo> pageFuzzyselectFive(Timestamp min,
			Timestamp max, Page<StayRegisterPo> vo) {
		int start=0;
		if (vo.getCurrentPage()>1) {
			start=(vo.getCurrentPage()-1)*vo.getPageSize();
		}
		List<StayRegisterPo> list= stayRegisterMapper.pageFuzzyselectFive(min, max,start, vo.getPageSize());
		vo.setResult(list);
		int count= stayRegisterMapper.countFuzzyselectFive(min, max);
		vo.setTotal(count);
		return vo;
	}

	@Override
	public List<StayRegisterPo> selectPayJingJianBanNot() {
		return this.stayRegisterMapper.selectPayJingJianBanNot();
	}

	@Override
	public List<StayRegisterPo> selectPayStayNumberNot() {
		return this.stayRegisterMapper.selectPayStayNumberNot();
	}

	@Override
	public List<StayRegisterPo> selectPayXiaoFeiNot() {
		return this.stayRegisterMapper.selectPayXiaoFeiNot();
	}

	@Override
	public List<StayRegisterPo> selectPayJingJianBan(Timestamp min,
			Timestamp max) {
		return this.stayRegisterMapper.selectPayJingJianBan(min, max);
	}

	@Override
	public List<StayRegisterPo> selectPayStayNumber(Timestamp min, Timestamp max) {
		return this.stayRegisterMapper.selectPayStayNumber(min, max);
	}

	@Override
	public List<StayRegisterPo> selectPayXiaoFei(Timestamp min, Timestamp max) {
		return this.stayRegisterMapper.selectPayXiaoFei(min, max);
	}

	@Override
	public List<StayRegisterPo> selectAllInformation(int id) {
		return this.stayRegisterMapper.selectAllInformation(id);
	}

	@Override
	public List<StayRegisterPo> selectXiaoFeiMingXi(int id) {
		return this.stayRegisterMapper.selectXiaoFeiMingXi(id);
	}

	@Override
	public int updateStayRegisterPredetermineID(Integer predetermineID,
			Integer id) {
		return this.stayRegisterMapper.updateStayRegisterPredetermineID(predetermineID, id);
	}

	@Override
	public StayRegisterPo selectSumconst(int id) {
		return this.stayRegisterMapper.selectSumconst(id);
	}



	

	

}
