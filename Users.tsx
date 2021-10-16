import React, { useEffect } from "react";

import { Link as RouterLink } from "react-router-dom";

import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";

import Typography from "@material-ui/core/Typography";

import Button from "@material-ui/core/Button";

import Container from "@material-ui/core/Container";

import Paper from "@material-ui/core/Paper";

import Box from "@material-ui/core/Box";

import Table from "@material-ui/core/Table";

import TableBody from "@material-ui/core/TableBody";

import TableCell from "@material-ui/core/TableCell";

import TableContainer from "@material-ui/core/TableContainer";

import TableHead from "@material-ui/core/TableHead";

import TableRow from "@material-ui/core/TableRow";

import { UsersInterface } from "../models/IUsers2";


import moment from 'moment';


 

const useStyles = makeStyles((theme: Theme) =>

 createStyles({

   container: {marginTop: theme.spacing(2)},

   table: { minWidth: 650},

   tableSpace: {marginTop: 20},

 })

);

 

function Users() {

 const classes = useStyles();

 const [pharmacist, setPharmas] = React.useState<UsersInterface[]>([]);

 

 const getPharmas = async () => {

   const apiUrl = "http://localhost:8080/users";

   const requestOptions = {

     method: "GET",

     headers: { "Content-Type": "application/json" },

   };

 

   fetch(apiUrl, requestOptions)
     .then((response) => response.json())
     .then((res) => {

       console.log(res.data);

       if (res.data) {

         setPharmas(res.data);

       } else {

         console.log("else");

       }

     });

 };

 

 useEffect(() => {

   getPharmas();

 }, []);

 

 return (

   <div>

     <Container className={classes.container} maxWidth="md">

       <Box display="flex">

         <Box flexGrow={1}>

           <Typography

             component="h2"

             variant="h6"

             color="primary"

             gutterBottom

           >

             ประวัติการจ่ายยา

           </Typography>

         </Box>

         <Box>

           <Button

             component={RouterLink}

             to="/create"

             variant="contained"

             color="primary"

           >

             บันทึกประวัติการจ่ายยา

           </Button>

         </Box>

       </Box>

       <TableContainer component={Paper} className={classes.tableSpace}>

         <Table className={classes.table} aria-label="simple table">

           <TableHead>

             <TableRow>

               <TableCell align="center" width="5%">

                 ID

               </TableCell>

               <TableCell align="center" width="20%">

                 ผู้บันทึก

               </TableCell>

               <TableCell align="center" width="25%">

                 ใบจ่ายยา

               </TableCell>

               <TableCell align="center" width="15%">

                 ชื่อผู้ป่วย

               </TableCell>

               <TableCell align="center" width="15%">

                 ยาที่จ่าย

               </TableCell>

               <TableCell align="center" width="20%">

                 วันที่และเวลา

               </TableCell>

             </TableRow>

           </TableHead>

           <TableBody>

             {pharmacist.map((user: UsersInterface) => (

               <TableRow key={user.ID}>

                 <TableCell align="right">{user.ID}</TableCell>

                 <TableCell align="left" size="medium">

                   {user.Treatment}

                 </TableCell>

                 <TableCell align="left">{user.Medicine}</TableCell>

                 <TableCell align="center">{moment(user.RecordTime).format("DD/MM/YYYY")}</TableCell>

               </TableRow>

             ))}

           </TableBody>

         </Table>

       </TableContainer>

     </Container>

   </div>

 );

}

 

export default Users;

