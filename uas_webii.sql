-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: May 24, 2024 at 09:40 AM
-- Server version: 10.4.28-MariaDB
-- PHP Version: 8.2.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `uas_webii`
--

-- --------------------------------------------------------

--
-- Table structure for table `ds_wisuda_tibil`
--

CREATE TABLE `ds_wisuda_tibil` (
  `nama` varchar(11) DEFAULT NULL,
  `nim` int(4) NOT NULL,
  `tmptlahir` varchar(12) DEFAULT NULL,
  `tgllahir` date DEFAULT NULL,
  `ipk` decimal(3,2) DEFAULT NULL,
  `predikat` varchar(17) DEFAULT NULL,
  `prodi` varchar(28) DEFAULT NULL,
  `masastudi` varchar(16) DEFAULT NULL,
  `suliet` int(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `ds_wisuda_tibil`
--

INSERT INTO `ds_wisuda_tibil` (`nama`, `nim`, `tmptlahir`, `tgllahir`, `ipk`, `predikat`, `prodi`, `masastudi`, `suliet`) VALUES
('mahasiswa1', 1001, 'Gumawang', '1999-12-10', 3.65, 'Sangat Memuaskan', 'Teknik Informatika Bilingual', '4 Tahun 6 Bulan', 533),
('mahasiswa2', 1002, 'Muara Bulian', '1999-08-24', 3.30, 'Sangat Memuaskan', 'Teknik Informatika Bilingual', '4 Tahun 6 Bulan', 500),
('mahasiswa3', 1003, 'Palembang', '1999-05-25', 3.42, 'Sangat Memuaskan', 'Teknik Informatika Bilingual', '4 Tahun 6 Bulan', 503),
('mahasiswa4', 1004, 'Palembang', '2000-01-03', 3.58, 'Sangat Memuaskan', 'Teknik Informatika Bilingual', '4 Tahun 6 Bulan', 507),
('mahasiswa5', 1005, 'Prabumulih', '1999-10-12', 3.40, 'Sangat Memuaskan', 'Teknik Informatika Bilingual', '4 Tahun 6 Bulan', 557),
('mahasiswa6', 1006, 'Palembang', '1999-05-09', 3.36, 'Sangat Memuaskan', 'Teknik Informatika Bilingual', '4 Tahun 6 Bulan', 513),
('mahasiswa7', 1007, 'Kembang Seri', '1999-10-09', 3.50, 'Sangat Memuaskan', 'Teknik Informatika Bilingual', '4 Tahun 6 Bulan', 513),
('mahasiswa8', 1008, 'Palembang', '2000-04-08', 3.99, 'Dengan Pujian', 'Teknik Informatika Bilingual', '3 Tahun 6 Bulan', 530),
('mahasiswa9', 1009, 'Palembang', '2000-12-20', 3.92, 'Dengan Pujian', 'Teknik Informatika Bilingual', '3 Tahun 6 Bulan', 530),
('mahasiswa10', 1010, 'Palembang', '2001-07-02', 3.88, 'Dengan Pujian', 'Teknik Informatika Bilingual', '3 Tahun 6 Bulan', 570),
('mahasiswa11', 1011, 'Palembang', '1999-06-25', 3.41, 'Sangat Memuaskan', 'Teknik Informatika Bilingual', '4 Tahun 7 Bulan', 507),
('mahasiswa12', 1012, 'Palembang', '1997-12-30', 3.07, 'Sangat Memuaskan', 'Teknik Informatika Bilingual', '4 Tahun 7 Bulan', 503),
('mahasiswa13', 1013, 'Palembang', '2000-01-05', 3.26, 'Sangat Memuaskan', 'Teknik Informatika Bilingual', '4 Tahun 7 Bulan', 507),
('mahasiswa14', 1014, 'Palembang', '1998-07-01', 3.27, 'Sangat Memuaskan', 'Teknik Informatika Bilingual', '4 Tahun 7 Bulan', 520),
('mahasiswa15', 1015, 'Palembang', '1999-05-21', 3.21, 'Sangat Memuaskan', 'Teknik Informatika Bilingual', '4 Tahun 7 Bulan', 500),
('mahasiswa16', 1016, 'Palembang', '1999-05-08', 3.82, 'Sangat Memuaskan', 'Teknik Informatika Bilingual', '4 Tahun 7 Bulan', 527),
('mahasiswa17', 1017, 'Prabumulih', '2000-01-01', 3.26, 'Sangat Memuaskan', 'Teknik Informatika Bilingual', '4 Tahun 8 Bulan', 580),
('mahasiswa18', 1018, 'Palembang', '1999-03-26', 3.28, 'Sangat Memuaskan', 'Teknik Informatika Bilingual', '4 Tahun 8 Bulan', 503),
('mahasiswa19', 1019, 'Palembang', '2001-01-31', 3.78, 'Dengan Pujian', 'Teknik Informatika Bilingual', '3 Tahun 8 Bulan', 537),
('mahasiswa20', 1020, 'Palembang', '1999-12-25', 3.33, 'Sangat Memuaskan', 'Teknik Informatika Bilingual', '4 Tahun 10 Bulan', 503),
('mahasiswa21', 1021, 'Palembang', '2000-07-04', 3.35, 'Sangat Memuaskan', 'Teknik Informatika Bilingual', '4 Tahun 10 Bulan', 517),
('mahasiswa22', 1022, 'Palembang', '2000-09-24', 3.90, 'dengan pujian', 'Teknik Informatika Bilingual', '3 Tahun 11 Bulan', 533),
('mahasiswa23', 1023, 'Palembang ', '1998-08-15', 3.01, 'Sangat Memuaskan ', 'Teknik Informatika Bilingual', '6 Tahun 0 Bulan', 417),
('mahasiswa24', 1024, 'Jakarta ', '1997-05-25', 3.41, 'Sangat Memuaskan ', 'Teknik Informatika Bilingual', '6 Tahun 0 Bulan', 503),
('mahasiswa25', 1025, 'Palembang ', '2000-02-21', 3.10, 'Sangat Memuaskan ', 'Teknik Informatika Bilingual', '5 Tahun 0 Bulan', 500),
('mahasiswa26', 1026, 'Palembang ', '1998-11-13', 3.28, 'Sangat Memuaskan ', 'Teknik Informatika Bilingual', '5 Tahun 0 Bulan', 500),
('mahasiswa27', 1027, 'Batam', '1999-05-04', 3.15, 'Sangat Memuaskan ', 'Teknik Informatika Bilingual', '5 Tahun 0 Bulan', 563),
('mahasiswa28', 1028, 'Lahat', '1999-11-28', 3.35, 'Sangat Memuaskan ', 'Teknik Informatika Bilingual', '5 Tahun 0 Bulan', 580),
('mahasiswa29', 1029, 'Prabumulih', '1999-04-11', 3.56, 'Sangat Memuaskan ', 'Teknik Informatika Bilingual', '5 Tahun 0 Bulan', 500),
('mahasiswa30', 1030, 'Belitang', '2000-08-15', 3.22, 'Sangat Memuaskan ', 'Teknik Informatika Bilingual', '5 Tahun 0 Bulan', 500),
('mahasiswa31', 1031, 'Jambi', '1999-04-18', 3.20, 'Sangat Memuaskan ', 'Teknik Informatika Bilingual', '5 Tahun 0 Bulan', 547),
('mahasiswa32', 1032, 'Jambi', '1999-09-16', 3.06, 'Sangat Memuaskan ', 'Teknik Informatika Bilingual', '5 Tahun 0 Bulan', 500),
('mahasiswa33', 1033, 'Tanjung Enim', '1999-08-18', 3.34, 'Sangat Memuaskan ', 'Teknik Informatika Bilingual', '5 Tahun 0 Bulan', 510),
('mahasiswa34', 1034, 'Muara Enim', '2000-01-12', 3.28, 'Sangat Memuaskan ', 'Teknik Informatika Bilingual', '5 Tahun 0 Bulan', 503),
('mahasiswa35', 1035, 'Air Balui', '1999-11-07', 3.16, 'Sangat Memuaskan ', 'Teknik Informatika Bilingual', '5 Tahun 0 Bulan', 500),
('mahasiswa36', 1036, 'Baturaja', '1999-12-01', 3.09, 'Sangat Memuaskan ', 'Teknik Informatika Bilingual', '5 Tahun 0 Bulan', 500),
('mahasiswa37', 1037, 'Baturaja', '1998-09-26', 3.28, 'Sangat Memuaskan ', 'Teknik Informatika Bilingual', '5 Tahun 0 Bulan', 517),
('mahasiswa38', 1038, 'Palembang ', '2000-03-20', 3.15, 'Sangat Memuaskan ', 'Teknik Informatika Bilingual', '5 Tahun 0 Bulan', 500),
('mahasiswa39', 1039, 'Palembang ', '1999-10-02', 3.15, 'Sangat Memuaskan ', 'Teknik Informatika Bilingual', '5 Tahun 0 Bulan', 500),
('mahasiswa40', 1040, 'Muara Enim', '2000-05-23', 3.40, 'Sangat Memuaskan ', 'Teknik Informatika Bilingual', '5 Tahun 0 Bulan', 507),
('mahasiswa41', 1041, 'Palembang ', '1999-06-22', 3.19, 'Sangat Memuaskan ', 'Teknik Informatika Bilingual', '5 Tahun 0 Bulan', 0),
('mahasiswa42', 1042, 'Lahat', '1999-04-15', 3.03, 'Sangat Memuaskan ', 'Teknik Informatika Bilingual', '5 Tahun 0 Bulan', 533),
('mahasiswa43', 1043, 'Palembang ', '1999-03-04', 3.29, 'Sangat Memuaskan ', 'Teknik Informatika Bilingual', '5 Tahun 0 Bulan', 583),
('mahasiswa44', 1044, 'Palembang ', '1999-10-18', 3.38, 'Sangat Memuaskan ', 'Teknik Informatika Bilingual', '5 Tahun 0 Bulan', 503),
('mahasiswa45', 1045, 'Palembang ', '2000-10-18', 3.38, 'Sangat Memuaskan ', 'Teknik Informatika Bilingual', '5 Tahun 0 Bulan', 500),
('mahasiswa46', 1046, 'Palembang ', '2000-03-23', 3.47, 'Sangat Memuaskan ', 'Teknik Informatika Bilingual', '5 Tahun 0 Bulan', 557),
('mahasiswa47', 1047, 'Palembang ', '1999-05-30', 2.96, 'Sangat Memuaskan ', 'Teknik Informatika Bilingual', '5 Tahun 0 Bulan', 567),
('mahasiswa48', 1048, 'Oku Timur', '1999-02-17', 3.01, 'Sangat Memuaskan ', 'Teknik Informatika Bilingual', '5 Tahun 0 Bulan', 500),
('mahasiswa49', 1049, 'Palembang ', '2000-12-09', 3.70, 'Dengan Pujian', 'Teknik Informatika Bilingual', '4 Tahun 0 Bulan', 510),
('mahasiswa50', 1050, 'Palembang ', '2001-06-13', 3.64, 'Dengan Pujian', 'Teknik Informatika Bilingual', '4 Tahun 0 Bulan', 500),
('mahasiswa51', 1051, 'Palembang ', '2000-05-07', 3.86, 'Dengan Pujian', 'Teknik Informatika Bilingual', '4 Tahun 0 Bulan', 553),
('mahasiswa52', 1052, 'Lubuklinggau', '2000-09-24', 3.63, 'Dengan Pujian', 'Teknik Informatika Bilingual', '4 Tahun 0 Bulan', 550),
('mahasiswa53', 1053, 'Nganjuk', '2000-04-17', 3.86, 'Dengan Pujian', 'Teknik Informatika Bilingual', '4 Tahun 0 Bulan', 507),
('mahasiswa54', 1054, 'Bengkulu', '1999-12-08', 3.56, 'Dengan Pujian', 'Teknik Informatika Bilingual', '4 Tahun 0 Bulan', 537),
('mahasiswa55', 1055, 'Jakarta ', '2000-02-05', 3.56, 'Dengan Pujian', 'Teknik Informatika Bilingual', '4 Tahun 0 Bulan', 547),
('mahasiswa56', 1056, 'Palembang ', '2001-01-07', 3.83, 'Dengan Pujian', 'Teknik Informatika Bilingual', '4 Tahun 0 Bulan', 597),
('mahasiswa57', 1057, 'Lawang Agung', '2000-10-28', 3.56, 'Sangat Memuaskan ', 'Teknik Informatika Bilingual', '4 Tahun 0 Bulan', 587),
('mahasiswa58', 1058, 'Lahat', '1998-10-20', 3.45, 'Sangat Memuaskan', 'Teknik Informatika Bilingual', '4 Tahun 2 Bulan', 510),
('mahasiswa59', 1059, 'Baturaja', '2000-05-21', 3.78, 'Sangat Memuaskan', 'Teknik Informatika Bilingual', '4 Tahun 2 Bulan', 500),
('mahasiswa60', 1060, 'Palembang', '2000-02-11', 3.40, 'Sangat Memuaskan', 'Teknik Informatika Bilingual', '4 Tahun 2 Bulan', 500),
('mahasiswa61', 1061, 'Palembang', '2001-05-21', 3.44, 'Sangat Memuaskan', 'Teknik Informatika Bilingual', '4 Tahun 2 Bulan', 530),
('mahasiswa62', 1062, 'Palembang', '2000-10-03', 3.58, 'Sangat Memuaskan', 'Teknik Informatika Bilingual', '4 Tahun 4 Bulan', 573),
('mahasiswa63', 1063, 'Bengkulu', '2000-07-30', 3.83, 'Sangat Memuaskan', 'Teknik Informatika Bilingual', '4 Tahun 4 Bulan', 580),
('mahasiswa64', 1064, 'Palembang', '1997-01-12', 3.38, 'Sangat Memuaskan', 'Teknik Informatika Bilingual', '5 Tahun 5 Bulan', 533),
('mahasiswa65', 1065, 'Palembang', '1999-01-21', 3.15, 'Sangat Memuaskan', 'Teknik Informatika Bilingual', '5 Tahun 5 Bulan', 530),
('mahasiswa66', 1066, 'Gumawang', '1999-12-10', 2.97, 'Sangat Memuaskan', 'Teknik Informatika Bilingual', '5 Tahun 5 Bulan', 553),
('mahasiswa67', 1067, 'Bukittinggi', '1999-12-27', 3.51, 'Sangat Memuaskan', 'Teknik Informatika Bilingual', '5 Tahun 5 Bulan', 447),
('mahasiswa68', 1068, 'Palembang', '1999-08-24', 3.16, 'Sangat Memuaskan', 'Teknik Informatika Bilingual', '5 Tahun 5 Bulan', 500),
('mahasiswa69', 1069, 'Palembang', '1998-12-03', 3.31, 'Sangat Memuaskan', 'Teknik Informatika Bilingual', '5 Tahun 5 Bulan', 500),
('mahasiswa70', 1070, 'Muaratawi', '1999-11-02', 3.18, 'Sangat Memuaskan', 'Teknik Informatika Bilingual', '5 Tahun 5 Bulan', 517),
('mahasiswa71', 1071, 'Palembang', '1999-03-15', 3.27, 'Sangat Memuaskan', 'Teknik Informatika Bilingual', '5 Tahun 5 Bulan', 473),
('mahasiswa72', 1072, 'Jambi', '1999-09-13', 3.44, 'Sangat Memuaskan', 'Teknik Informatika Bilingual', '5 Tahun 5 Bulan', 543),
('mahasiswa73', 1073, 'Palembang', '1999-06-08', 3.23, 'Sangat Memuaskan', 'Teknik Informatika Bilingual', '5 Tahun 5 Bulan', 513),
('mahasiswa74', 1074, 'Muara Enim', '2000-11-30', 3.27, 'Sangat Memuaskan', 'Teknik Informatika Bilingual', '4 Tahun 5 Bulan', 500),
('mahasiswa75', 1075, 'Belitang', '1999-05-12', 3.23, 'Sangat Memuaskan', 'Teknik Informatika Bilingual', '5 Tahun 5 Bulan', 503);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `ds_wisuda_tibil`
--
ALTER TABLE `ds_wisuda_tibil`
  ADD PRIMARY KEY (`nim`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
