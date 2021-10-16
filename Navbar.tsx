import React from "react";

import { makeStyles } from "@material-ui/core/styles";

import AppBar from "@material-ui/core/AppBar";

import Toolbar from "@material-ui/core/Toolbar";

import Typography from "@material-ui/core/Typography";

import IconButton from "@material-ui/core/IconButton";

import MenuIcon from "@material-ui/icons/Menu";

import { Link } from "react-router-dom";

import Avatar from '@material-ui/core/Avatar';
import { Button } from "@material-ui/core";
import { createTheme } from '@material-ui/core/styles';

const useStyles = makeStyles((theme) => ({
  
  nav : {background : "#239B56"},
  root: {flexGrow: 2},

 menuButton: {marginRight: theme.spacing(2)},
 
 title: {flexGrow: 1},
 large: {
    width: theme.spacing(7),
    height: theme.spacing(7),
  },
 
 navlink: {color: "white",textDecoration: "none"},

}));


function Navbar() {

 const classes = useStyles();

 return (

   <div className={classes.root} >
   <AppBar position="static" className={classes.nav}>
        
        <Toolbar >
          <IconButton edge="start" className={classes.menuButton} color="inherit" aria-label="menu">
          <img src="/img/med-jar.png" width="50px"></img>
            {/*<MenuIcon />*/}
          </IconButton>
          {/*<Link className={classes.navlink} to="/">*/}
          <Typography className={classes.title}>
             <h2>ระบบบันทึกการจ่ายยาคนไข้ใน</h2>
          </Typography>
         {/* </Link>*/}
         
         <Avatar style={{ float: "right" }} alt="Chawarat Narit" src="/img/Profile.jpg" className={classes.large} /><br/>
         &nbsp;&nbsp;
         <p>Chawarat Narit</p><br/>
         &nbsp;&nbsp;
         <Button

             style={{ float: "right" ,backgroundColor : "#F5B041" }}
              
             

             variant="contained"

             color="primary"

           >Logout
        </Button>
        </Toolbar>
        
      </AppBar>
     
   </div>

 );

}

export default Navbar;