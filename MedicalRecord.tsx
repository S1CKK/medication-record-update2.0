import React, { useDebugValue, useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import { makeStyles, Theme, createStyles } from "@material-ui/core/styles";
import TextField from "@material-ui/core/TextField";
import Button from "@material-ui/core/Button";
import FormControl from "@material-ui/core/FormControl";
import Container from "@material-ui/core/Container";
import Paper from "@material-ui/core/Paper";
import Grid from "@material-ui/core/Grid";
import Box from "@material-ui/core/Box";
import Typography from "@material-ui/core/Typography";
import Divider from "@material-ui/core/Divider";
import Snackbar from "@material-ui/core/Snackbar";
import MuiAlert, { AlertProps } from "@material-ui/lab/Alert";
import { UsersInterface } from "../models/IUsers2";
import {MuiPickersUtilsProvider,KeyboardDatePicker,} from "@material-ui/pickers";
import DateFnsUtils from "@date-io/date-fns";
import Avatar from '@material-ui/core/Avatar';
import InputLabel from '@material-ui/core/InputLabel';
import FormHelperText from '@material-ui/core/FormHelperText';
import Select from '@material-ui/core/Select';
import NativeSelect from '@material-ui/core/NativeSelect';
import {PharmacistInterface} from '../models/IPharmas';
import {TreatmentInterface} from '../models/ITreatment';
import {AdmissionInterface} from '../models/IAdmission';
import {MedicineInterface} from '../models/IMedicine';

function Alert(props: AlertProps) {
 return <MuiAlert elevation={6} variant="filled" {...props} />;
}

const useStyles = makeStyles((theme: Theme) =>  

 createStyles({
   root: {flexGrow: 1},
   container: {marginTop: theme.spacing(2)},
   paper: {padding: theme.spacing(2),color: theme.palette.text.secondary},   
   formControl: {
    margin: theme.spacing(1),
    minWidth: 120,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
 })
);

function UserCreate() {
    const [name, setName] = React.useState({});
    const handleChange1 = (event:React.ChangeEvent<{ name?: string; value: unknown }>) => {
        setName(event.target.value as number);
      };
    const [order_id, setOrder_id] = React.useState({});
    const handleChange2 = (event:any) => {
        setOrder_id(event.target.value);
      };
    const [an_id, setAN_id] = React.useState({});
    const [an_name, setAN_name] = React.useState({});
    const [amount, setAmount] = React.useState({});
    const handleChange3 = (event:any) => {
        setAN_name(event.target.value);
      };
    const [med_name, setMed_name] = React.useState({});
    const handleChange4 = (event:any) => {
        setMed_name(event.target.value);   
      };
    const classes = useStyles();
    const [selectedDate, setSelectedDate] = React.useState<Date | null>(
    new Date()
 );

 const [user, setUser] = React.useState<Partial<UsersInterface>>({});
 const [success, setSuccess] = React.useState(false);
 const [error, setError] = React.useState(false);
 const handleClose = (event?: React.SyntheticEvent, reason?: string) => {
   if (reason === "clickaway") {
     return;
   }
   setSuccess(false);
   setError(false);
 };

  const handleDateChange = (date: Date | null) => {
   setSelectedDate(date);
 };

  const handleInputChange = (
   event: React.ChangeEvent<{ id?: string; value: any }>
 ) => {

   const id = event.target.id as keyof typeof UserCreate;
   const { value } = event.target;
   setUser({ ...user, [id]: value });
 };

 const [ pharmacists, setPharmacists ] = React.useState<PharmacistInterface[]>([]);
 const [ treatments, setTreatments ] = React.useState<TreatmentInterface[]>([]);
 const [ admissions, setAdmissions ] = React.useState<AdmissionInterface[]>([]);
 const [ medications, setMedicines ] = React.useState<MedicineInterface[]>([]);
  

  const getPharmacists = async () => {
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
        setPharmacists(res.data);
      } else {
        console.log("else");
      }
    });
 }
 
 const getTreatments = async () => {
  const apiUrl = "http://localhost:8080/treatments";

  const requestOptions = {
    method: "GET",
    headers: { "Content-Type": "application/json" },
  };

  fetch(apiUrl, requestOptions)
    .then((response) => response.json())
    .then((res) => {
      console.log(res.data);
      if (res.data) {
        setTreatments(res.data);
      } else {
        console.log("else");
      }
    });
 }

 const getAdmissions = async () => {
  const apiUrl = "http://localhost:8080/admissions";

  const requestOptions = {
    method: "GET",
    headers: { "Content-Type": "application/json" },
  };

  fetch(apiUrl, requestOptions)
    .then((response) => response.json())
    .then((res) => {
      console.log(res.data);
      if (res.data) {
        setAdmissions(res.data);
      } else {
        console.log("else");
      }
    });
 }

 const getMedicines = async () => {
  const apiUrl = "http://localhost:8080/medicines";

  const requestOptions = {
    method: "GET",
    headers: { "Content-Type": "application/json" },
  };

  fetch(apiUrl, requestOptions)
    .then((response) => response.json())
    .then((res) => {
      console.log(res.data);
      if (res.data) {
        setMedicines(res.data);
      } else {
        console.log("else");
      }
    });
 }

 function submit() {

   let data = {
     PharmaID: name?? "",
   /*  Treatments: order_id?? "",
     Medicine: med_name ?? "",
     Amount: user.Amount ?? "",
     RecordTime: selectedDate,*/
   };

   const apiUrl = "http://localhost:8080/medicationrecords";
   const requestOptions = {
     method: "POST",
     headers: { "Content-Type": "application/json" },
     body: JSON.stringify(data),
   };

   fetch(apiUrl, requestOptions)
     .then((response) => response.json())
     .then((res) => {
       if (res.data) {
         setSuccess(true);
       } else {
         setError(true);
       }
     });    
 }
 useEffect(() => {

  getPharmacists();
  getTreatments();
  getAdmissions();
  getMedicines();

}, []);
 return (
   
   <Container className={classes.container} maxWidth="md">
     <Snackbar open={success} autoHideDuration={6000} onClose={handleClose}>
       <Alert onClose={handleClose} severity="success">
         บันทึกข้อมูลสำเร็จ
       </Alert>
     </Snackbar>
     <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
       <Alert onClose={handleClose} severity="error">
         บันทึกข้อมูลไม่สำเร็จ
       </Alert>
     </Snackbar>
     <Paper className={classes.paper}>
       <Box display="flex">
         <Box flexGrow={1}>
           <Typography
             component="h2"
             variant="h6"
             color="primary"
             gutterBottom
           >
             กรุณากรอกข้อมูล
           </Typography>
         </Box>
       </Box>
       <Divider />   
       <Grid container spacing={3} className={classes.root}>
         <Grid item xs={6}>
           <p>ผู้บันทึก</p>
        <FormControl fullWidth variant="outlined" className={classes.formControl}>
        <InputLabel htmlFor="outlined-age-native-simple"></InputLabel>
        <Select        
          native
          value={name}
          onChange={handleChange1}
          label="Name"
          inputProps={{
            id: 'name',
          }}
        >
          <option aria-label="None" value="" />
          { pharmacists.map( (user:PharmacistInterface) => (
              <option value={user.ID}>{user.Name}</option>
          ) ) }
         
        </Select>
      </FormControl>
         </Grid>
        <Grid item xs={6}>
           <p>ใบบันทึกการรักษา</p>
        <FormControl fullWidth variant="outlined" className={classes.formControl}>
        <InputLabel htmlFor="outlined-age-native-simple"></InputLabel>
        <Select
          native
          value={order_id}
          onChange={handleChange2}
          label="Order_id"
          inputProps={{
            id: 'order_id',
          }}
        >
          <option aria-label="None" value="" />
          { treatments.map( (admit:TreatmentInterface) => (
              <option value={admit.ID}> {admit.ID} </option>
          ) ) }
 
        </Select>
      </FormControl>
         </Grid>

        <Grid item xs={6}>
           <p>ยาที่จ่าย</p>
        <FormControl fullWidth variant="outlined" className={classes.formControl}>
        <InputLabel htmlFor="outlined-age-native-simple"></InputLabel>
        <Select
          native
          value={med_name}
          onChange={handleChange4}
          label="Med_name"
          inputProps={{
            id: 'med_name',
          }}
        >
          <option aria-label="None" value="" />
          { medications.map( (med:MedicineInterface) => (
              <option value={med.ID}> {med.Name}</option>
          ) ) }
        </Select>
      </FormControl>
         </Grid>
         <Grid item xs={6}>
          <FormControl fullWidth variant="outlined" className={classes.formControl}>
          <p>จำนวน</p>
          <TextField
            id="Amount"
            variant="outlined"
            type="number"
            size="medium"
            value={user.Amount || ""}
            onChange={handleInputChange}
            />
          </FormControl>
          </Grid>          
         <Grid item xs={6}>
           <FormControl fullWidth variant="outlined">
             <p>วันที่และเวลา</p>
             <MuiPickersUtilsProvider utils={DateFnsUtils}>
               <KeyboardDatePicker
                 margin="normal"
                 id="RecordDate"
                 format="yyyy-MM-dd"
                 value={selectedDate}
                 onChange={handleDateChange}
                 KeyboardButtonProps={{
                   "aria-label": "change date",
                 }}
               />
             </MuiPickersUtilsProvider>
           </FormControl>
         </Grid>
         <Grid item xs={12}>
           <Button component={RouterLink} to="/" variant="contained">
             ย้อนกลับ
           </Button>
           <Button
             style={{ float: "right" ,backgroundColor : "#239B56"}}
             onClick={submit}
             variant="contained"
             color="primary"
           >
             บันทึกข้อมูล
           </Button>
         </Grid>
       </Grid>
     </Paper>
   </Container>   
 );
}
export default UserCreate;