import React, { FC, useEffect } from 'react';
import Avatar from '@material-ui/core/Avatar';
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import TextField from '@material-ui/core/TextField';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Checkbox from '@material-ui/core/Checkbox';
import Link from '@material-ui/core/Link';
import Paper from '@material-ui/core/Paper';
import Box from '@material-ui/core/Box';
import Grid from '@material-ui/core/Grid';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import Swal from 'sweetalert2'; // alert



function Copyright() {
  return (
    <Typography variant="body2" color="textSecondary" align="center">
      {'Copyright © newler55 | '}
      <Link color="inherit" href="https://github.com/newler55">
        Your Github
      </Link>{' '}
      {new Date().getFullYear()}
      {'.'}
    </Typography>
  );
}

const useStyles = makeStyles((theme) => ({
  root: {
    height: '98vh',
  },
  image: {
    backgroundImage: 'url(https://scontent.fnak3-1.fna.fbcdn.net/v/t1.0-9/50408628_1024226204452974_6430172104947138560_o.jpg?_nc_cat=100&ccb=2&_nc_sid=dd9801&_nc_ohc=EiGvlsrgSKMAX82jfoY&_nc_ht=scontent.fnak3-1.fna&oh=9ba8a2a068fb52af16f06794dc8de422&oe=5FFF7EF5)',
    // backgroundImage: {bgLogin},
    backgroundRepeat: 'no-repeat',
    backgroundColor:
      theme.palette.type === 'light' ? theme.palette.grey[50] : theme.palette.grey[900],
    backgroundSize: 'cover',
    backgroundPosition: 'center',
  },
  paper: {
    margin: theme.spacing(8, 4),
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
  },
  large: {
    width: 100,
    height: 100,
  },
  form: {
    width: '100%', // Fix IE 11 issue.
    marginTop: theme.spacing(1),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
}));

interface Login {
  username: string;
  password: string;

}

const Login: FC<{}> = () => {
  const classes = useStyles();

  const [login, setLogin] = React.useState<Partial<Login>>({});

  // alert setting
  const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
    didOpen: toast => {
      toast.addEventListener('mouseenter', Swal.stopTimer);
      toast.addEventListener('mouseleave', Swal.resumeTimer);
    },
  });


  function redirecLogin() {
    if ((login.username == "B6015695" && login.password == "1234") ||
      (login.username == "admin" && login.password == "1234")
    ) {
      Toast.fire({
        icon: 'success',
        title: 'เข้าสู่ระบบสำเร็จ',
      });
      //redirec Page ... http://localhost:3000/Table
      window.location.href = "http://localhost:3000/Table";
      console.log("LOGIN TO DISEASE");
    } else {
      Toast.fire({
        icon: 'error',
        title: 'username หรือ password ไม่ถูกต้อง',
      });
    }

  }

  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof login;
    const { value } = event.target;
    setLogin({ ...login, [name]: value });
    console.log(login);
  };

  return (
    <Grid container component="main" className={classes.root}>
      <CssBaseline />
      <Grid item xs={false} sm={4} md={7} className={classes.image} />
      <Grid item xs={12} sm={8} md={5} component={Paper} elevation={6} square>
        <div className={classes.paper}>
          <Avatar src="https://pbs.twimg.com/profile_images/1141646207837564928/nhOyjcxQ_400x400.png" className={classes.large} />
          <Typography component="h1" variant="h2">
            เข้าสู่ระบบ
          </Typography>
          <form className={classes.form} noValidate>
            <TextField
              variant="outlined"
              margin="normal"
              required
              fullWidth
              id="username"
              label="ชื่อผู้ใช้"
              name="username"
              autoComplete="username"
              autoFocus
              value={login.username || ""}
              onChange={handleChange}

            />
            <TextField
              variant="outlined"
              margin="normal"
              required
              fullWidth
              name="password"
              label="รหัสผ่าน"
              type="password"
              id="password"
              autoComplete="current-password"
              value={login.password || ""}
              onChange={handleChange}

            />
            <FormControlLabel
              control={<Checkbox value="remember" color="primary" />}
              label=" จดจำฉัน"
            />
            <Button
              fullWidth
              variant="contained"
              color="primary"
              onClick={redirecLogin}
            >
              เข้าสู่ระบบ
            </Button>
            <Grid container>
              <Grid item xs>
                <Link href="#" variant="body2">
                  ลืมรหัสผ่าน!?
                </Link>
              </Grid>
              <Grid item>
                <Link href="#" variant="body2">
                  {"ยังไม่มีบัญชี? สมัครสมาชิก"}
                </Link>
              </Grid>
            </Grid>
            <Box mt={5}>
              <Copyright />
            </Box>
          </form>
        </div>
      </Grid>
    </Grid>
  );
};


export default Login;