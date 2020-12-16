import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import { DefaultApi } from '../../api/apis';
import MenuIcon from '@material-ui/icons/Menu';
import Avatar from '@material-ui/core/Avatar';
import Link from '@material-ui/core/Link';
import {
  Content,
  ContentHeader,
} from '@backstage/core';
import { Link as RouterLink } from 'react-router-dom';
import {
  Button,
  AppBar,
  Toolbar,
  Typography,
  IconButton
} from '@material-ui/core';


const useStyles = makeStyles(theme => ({
  table: {
    minWidth: 650,
  },
  title: {
    flexGrow: 1,
  },
  menuButton: {
    marginRight: theme.spacing(2),
  },
  small: {
    width: 42,
    height: 42,
  },
  logoutButton: {
    marginLeft: 10,
    marginRight: 10,
    color: 'white'
  }
}));

function redirecLogOut() {
  // redire Page ... http://localhost:3000/
  window.location.href = "http://localhost:3000/";
}

export default function ComponentsTable() {
  const classes = useStyles();
  const api = new DefaultApi();
  const [diseases, setDisease] = useState(Array);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const getDiseases = async () => {
      const res = await api.listDisease({ limit: 10, offset: 0 });
      setLoading(false);
      setDisease(res);
    };
    getDiseases();
  }, [loading]);



  const deleteDiseases = async (id: number) => {
    const res = await api.deleteDisease({ id: id });
    setLoading(true);
  };

  return (

    <div>
      <AppBar position="static">
        <Toolbar>
          <IconButton edge="start" className={classes.menuButton} color="inherit" aria-label="menu">
            <MenuIcon />
          </IconButton>
          <Typography variant="h4" className={classes.title}>
            ระบบบันทึกข้อมูลผู้ป่วยใน
          </Typography>
          <IconButton
            aria-label="account of current user"
            aria-controls="menu-appbar"
            aria-haspopup="true"
            color="inherit"
          >
            <Avatar alt="Remy Sha" src="https://scontent.fnak3-1.fna.fbcdn.net/v/t1.0-9/70458876_2383884038537083_7510417574585171968_n.jpg?_nc_cat=107&ccb=2&_nc_sid=174925&_nc_eui2=AeHrhWheLyHhKLvfx-GStXBs2pFiOaifHOvakWI5qJ8c6wGnrMqKjBPtNI93BRmPd8R3HMeup0e4VtorapXFRD5p&_nc_ohc=tm9sByGdWJ8AX8YDEQ1&_nc_ht=scontent.fnak3-1.fna&oh=aecbe134c66360c5d66d5f1b321dcef8&oe=5FFF1D26" className={classes.small} />
            <Typography>
              <Link variant="h6" onClick={redirecLogOut} className={classes.logoutButton}>
                ออกจากระบบ
                </Link>
            </Typography>
          </IconButton>
        </Toolbar>
      </AppBar>
      <Content>
        <ContentHeader title="">
          <Button
            size="large"
            style={{ float: 'right', marginBottom: 'auto' }}
            color="primary"
            component={RouterLink}
            to="/Disease"
            variant="contained"
          >
            เพิ่มข้อมูลผู้ป่วยใน
        </Button>
        </ContentHeader>

        <TableContainer component={Paper}>
          <Table className={classes.table} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="center">แพทย์เจ้าของไข้</TableCell>
                <TableCell align="center">ผู้ป่วย</TableCell>
                <TableCell align="center">ระดับความรุนเเรง</TableCell>
                <TableCell align="center">อาการ</TableCell>
                <TableCell align="center">วันเข้าพัก</TableCell>
                <TableCell align="center">ห้องพักฟื้น</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {diseases.map((item: any) => (
                <TableRow key={item.id}>

                  <TableCell align="center">{item.edges?.employee?.userId}</TableCell>
                  <TableCell align="center">{item.name}</TableCell>
                  <TableCell align="center">{item.edges?.severity?.name}</TableCell>
                  <TableCell align="center">{item.symptom}</TableCell>
                  <TableCell align="center">{item.contagion}</TableCell>
                  <TableCell align="center">{item.edges?.diseasetype?.name}</TableCell>
                  <TableCell align="center">
                    <Button
                      onClick={() => {
                        deleteDiseases(item.id);
                      }}
                      style={{ marginLeft: 10 }}
                      variant="contained"
                      color="secondary"
                    >
                      ลบข้อมูล
               </Button>
                  </TableCell>
                </TableRow>
              ))}

              {/* <TableRow >
                <TableCell align="center">นพ.เชาวนะ ดุสิตนานนท์</TableCell>
                <TableCell align="center">นาย บุญลือ วงษ์ท้าว</TableCell>
                <TableCell align="center">เริ่มต้น</TableCell>
                <TableCell align="center">ท้องเสีย</TableCell>
                <TableCell align="center">07/09/2020 10:20:00</TableCell>
                <TableCell align="center">F03291</TableCell>
                <TableCell align="center">
                  <Button
                    style={{ marginLeft: 10 }}
                    variant="contained"
                    color="secondary"
                  >
                    ลบข้อมูล
                  </Button>
                </TableCell>
              </TableRow>

              <TableRow >
                <TableCell align="center">พญ.รุ่งลัดดา ตัณฑวิเชียร</TableCell>
                <TableCell align="center">นาย ประเสริฐ อังกูรไกรวิชญ์</TableCell>
                <TableCell align="center">รุนเเรง</TableCell>
                <TableCell align="center">เส้นเลือดในสมองตีบ</TableCell>
                <TableCell align="center">09/10/2020 11:23:00</TableCell>
                <TableCell align="center">C48572</TableCell>
                <TableCell align="center">
                  <Button
                    style={{ marginLeft: 10 }}
                    variant="contained"
                    color="secondary"
                  >
                    ลบข้อมูล
                  </Button>
                </TableCell>
               </TableRow>

              <TableRow >
                <TableCell align="center">ทพญ.เบญจวรรณ เพิ่มไชยศิริ</TableCell>
                <TableCell align="center">นาย ธเรศ เล่าสุอังกูร</TableCell>
                <TableCell align="center">ปานกลาง</TableCell>
                <TableCell align="center">เบาหวาน</TableCell>
                <TableCell align="center">11/12/2020 13:45:00</TableCell>
                <TableCell align="center">F00984</TableCell>
                <TableCell align="center">
                  <Button
                    style={{ marginLeft: 10 }}
                    variant="contained"
                    color="secondary"
                  >
                    ลบข้อมูล
                  </Button>
                </TableCell>
              </TableRow> */}


            </TableBody>
          </Table>
        </TableContainer>
      </Content>
    </div>
  );
}
