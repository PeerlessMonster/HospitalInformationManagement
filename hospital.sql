-- phpMyAdmin SQL Dump
-- version 4.9.1
-- https://www.phpmyadmin.net/
--
-- 主机： localhost
-- 生成日期： 2023-05-24 15:23:55
-- 服务器版本： 8.0.17
-- PHP 版本： 7.3.10

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `hospital`
--

-- --------------------------------------------------------

--
-- 表的结构 `admin`
--

CREATE TABLE `admin` (
  `employeeNo` varchar(20) NOT NULL,
  `password` varchar(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- 转存表中的数据 `admin`
--

INSERT INTO `admin` (`employeeNo`, `password`) VALUES
('AD01', 'admin');

-- --------------------------------------------------------

--
-- 表的结构 `doctor`
--

CREATE TABLE `doctor` (
  `employeeNo` varchar(20) NOT NULL,
  `password` varchar(20) NOT NULL,
  `doctorName` varchar(20) NOT NULL,
  `sex` char(2) NOT NULL,
  `department` varchar(18) NOT NULL,
  `staff` varchar(10) NOT NULL,
  `doctorPhone` char(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- 转存表中的数据 `doctor`
--

INSERT INTO `doctor` (`employeeNo`, `password`, `doctorName`, `sex`, `department`, `staff`, `doctorPhone`) VALUES
('DC001', 'zhongnanshan', '钟南山', '男', '呼吸科', '主任', '12345678910'),
('DC002', 'lilanjuan', '李兰娟', '女', '传染病科', '主任', '98765432100'),
('DC003', 'liming', '李明', '男', '骨科', '主治医师', '18999999999');

-- --------------------------------------------------------

--
-- 表的结构 `expense`
--

CREATE TABLE `expense` (
  `staff` varchar(10) NOT NULL,
  `expense` decimal(5,2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- 转存表中的数据 `expense`
--

INSERT INTO `expense` (`staff`, `expense`) VALUES
('主任', '20.00'),
('副主任', '15.00'),
('主治医师', '8.00');

-- --------------------------------------------------------

--
-- 表的结构 `hospital_file`
--

CREATE TABLE `hospital_file` (
  `fileNo` varchar(15) NOT NULL,
  `patientNo` varchar(20) NOT NULL,
  `balance` decimal(10,2) NOT NULL,
  `inDate` date NOT NULL,
  `outDate` date DEFAULT NULL,
  `roomNo` varchar(4) NOT NULL,
  `bedNo` varchar(3) NOT NULL,
  `employeeNo` varchar(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- 转存表中的数据 `hospital_file`
--

INSERT INTO `hospital_file` (`fileNo`, `patientNo`, `balance`, `inDate`, `outDate`, `roomNo`, `bedNo`, `employeeNo`) VALUES
('PA0000004F01', 'PA0000004', '100000.00', '2023-05-23', NULL, 'R01', 'B1', 'DC003');

-- --------------------------------------------------------

--
-- 表的结构 `hospital_record`
--

CREATE TABLE `hospital_record` (
  `date` date NOT NULL,
  `hospitalRecordNo` varchar(15) NOT NULL,
  `fileNo` varchar(15) NOT NULL,
  `symptom` varchar(100) NOT NULL,
  `listNo` varchar(15) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- 转存表中的数据 `hospital_record`
--

INSERT INTO `hospital_record` (`date`, `hospitalRecordNo`, `fileNo`, `symptom`, `listNo`) VALUES
('2023-05-24', 'H20230524001', 'PA0000004F01', '观察下情况先，暂时不用药', NULL);

-- --------------------------------------------------------

--
-- 表的结构 `list`
--

CREATE TABLE `list` (
  `listNo` varchar(15) NOT NULL,
  `medicineNo` char(11) NOT NULL,
  `count` int(11) NOT NULL,
  `method` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- 转存表中的数据 `list`
--

INSERT INTO `list` (`listNo`, `medicineNo`, `count`, `method`) VALUES
('L20230515001', 'H12345678', 6, '口服，每日三次，每次三片'),
('L20230515001', 'H12345679', 1, '外用，每日一次'),
('L20230524001', 'H87654321', 2, '口服，每日一次，每次两片');

-- --------------------------------------------------------

--
-- 表的结构 `medicine`
--

CREATE TABLE `medicine` (
  `medicineNo` char(11) NOT NULL,
  `medicineName` varchar(30) NOT NULL,
  `medicinePrice` decimal(10,2) NOT NULL,
  `inventory` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- 转存表中的数据 `medicine`
--

INSERT INTO `medicine` (`medicineNo`, `medicineName`, `medicinePrice`, `inventory`) VALUES
('H12345678', '云南白药喷雾剂', '25.80', 66666),
('H12345679', '跌打损伤红花油', '30.60', 99999),
('H87654321', '莲花清瘟胶囊', '9.90', 2333);

-- --------------------------------------------------------

--
-- 表的结构 `patient`
--

CREATE TABLE `patient` (
  `patientNo` varchar(20) NOT NULL,
  `password` varchar(20) NOT NULL,
  `patientName` varchar(20) NOT NULL,
  `sex` char(2) NOT NULL,
  `patientPhone` char(11) NOT NULL,
  `patientAddress` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- 转存表中的数据 `patient`
--

INSERT INTO `patient` (`patientNo`, `password`, `patientName`, `sex`, `patientPhone`, `patientAddress`) VALUES
('PA0000002', 'meimei8000', '韩梅梅', '女', '13800138000', NULL),
('PA0000003', 'lei0831', '李雷', '男', '13800000831', '某某省某某市某某区某某路某某号'),
('PA0000004', 'mingshi0866', '无名氏', '男', '10086100866', NULL);

-- --------------------------------------------------------

--
-- 表的结构 `registrate`
--

CREATE TABLE `registrate` (
  `theOrder` int(11) NOT NULL,
  `patientNo` varchar(20) NOT NULL,
  `employeeNo` varchar(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- 转存表中的数据 `registrate`
--

INSERT INTO `registrate` (`theOrder`, `patientNo`, `employeeNo`) VALUES
(1, 'PA0000001', 'DC001'),
(2, 'PA0000003', 'DC001'),
(3, 'PA0000002', 'DC002');

-- --------------------------------------------------------

--
-- 表的结构 `room`
--

CREATE TABLE `room` (
  `roomNo` varchar(4) NOT NULL,
  `location` varchar(20) NOT NULL,
  `roomPrice` decimal(10,2) NOT NULL,
  `department` varchar(18) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- 转存表中的数据 `room`
--

INSERT INTO `room` (`roomNo`, `location`, `roomPrice`, `department`) VALUES
('R01', '住院部3楼', '60.00', '骨科');

-- --------------------------------------------------------

--
-- 表的结构 `visit_record`
--

CREATE TABLE `visit_record` (
  `date` date NOT NULL,
  `visitRecordNo` varchar(15) NOT NULL,
  `patientNo` varchar(20) NOT NULL,
  `employeeNo` varchar(20) NOT NULL,
  `symptom` varchar(100) NOT NULL,
  `listNo` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- 转存表中的数据 `visit_record`
--

INSERT INTO `visit_record` (`date`, `visitRecordNo`, `patientNo`, `employeeNo`, `symptom`, `listNo`) VALUES
('2023-05-15', 'V20230515001', 'PA0000004', 'DC002', '年轻人太气盛！', 'L20230515001'),
('2023-05-24', 'V20230524001', 'PA0000003', 'DC002', '二阳了……', 'L20230524001');

-- --------------------------------------------------------

--
-- 表的结构 `work`
--

CREATE TABLE `work` (
  `employeeNo` varchar(20) NOT NULL,
  `section` enum('门诊部','住院部') CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- 转存表中的数据 `work`
--

INSERT INTO `work` (`employeeNo`, `section`) VALUES
('DC001', '住院部'),
('DC002', '门诊部'),
('DC003', '住院部');

--
-- 转储表的索引
--

--
-- 表的索引 `admin`
--
ALTER TABLE `admin`
  ADD PRIMARY KEY (`employeeNo`);

--
-- 表的索引 `doctor`
--
ALTER TABLE `doctor`
  ADD PRIMARY KEY (`employeeNo`);

--
-- 表的索引 `hospital_file`
--
ALTER TABLE `hospital_file`
  ADD PRIMARY KEY (`fileNo`);

--
-- 表的索引 `hospital_record`
--
ALTER TABLE `hospital_record`
  ADD PRIMARY KEY (`hospitalRecordNo`);

--
-- 表的索引 `medicine`
--
ALTER TABLE `medicine`
  ADD PRIMARY KEY (`medicineNo`);

--
-- 表的索引 `patient`
--
ALTER TABLE `patient`
  ADD PRIMARY KEY (`patientNo`);

--
-- 表的索引 `registrate`
--
ALTER TABLE `registrate`
  ADD PRIMARY KEY (`theOrder`);

--
-- 表的索引 `room`
--
ALTER TABLE `room`
  ADD PRIMARY KEY (`roomNo`);

--
-- 表的索引 `visit_record`
--
ALTER TABLE `visit_record`
  ADD PRIMARY KEY (`visitRecordNo`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `registrate`
--
ALTER TABLE `registrate`
  MODIFY `theOrder` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
