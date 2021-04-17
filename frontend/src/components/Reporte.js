import {React, useState} from 'react'
import {Header, Segment, Icon, Container, Grid, Input, Button, Image, Dropdown} from 'semantic-ui-react'
import NavBar from './NavBar'
import '../css/Content.css'
import { useHistory } from 'react-router'
const axios = require('axios').default

function Reporte() {
    const history = useHistory()
    if (localStorage.getItem("LOGED") == null){
        history.push("/Login")
    }else if (localStorage.getItem("LOGED")=="Cliente"){
        history.push("/Home")
    }
    const [imagen, setImagen] = useState('')
    const [title, setTitle] = useState('')
    const [year, setYear] = useState('')
    const [month, setMonth] = useState('0')
    const [cat, setCat] = useState('')
    const [day, setDay] = useState('')
    const [cipher, setCipher] = useState('')
    function getMonth(m){
        switch (m) {
            case 'ENERO':
                return '1'
            case 'FEBRERO':
                return '2'
            case 'MARZO':
                return '3'
            case 'ABRIL':
                return '4'
            case 'MAYO':
                return '5'
            case 'JUNIO':
                return '6'
            case 'JULIO':
                return '7'
            case 'AGOSTO':
                return '8'
            case 'SEPTIEMBRE':
                return '9'
            case 'OCTUBRE':
                return '10'
            case 'NOVIEMBRE':
                return '11'
            case 'DICIEMBRE':
                return '12'
            default:
                return '0';
        }
    }
    
    function getCipher(m){
        switch (m) {
            case 'Sin Cifrar':
                return 0
            case 'Cifrado Sensible':
                return 1
            case 'Cifrado':
                return 2
            default:
                break;
        }
    }

    const dayOptions =[
        {key: 1, text: 1, value: 1},
        {key: 2, text: 2, value: 2},
        {key: 3, text: 3, value: 3},
        {key: 4, text: 4, value: 4},
        {key: 5, text: 5, value: 5},
        {key: 6, text: 6, value: 6},
        {key: 7, text: 7, value: 7},
        {key: 8, text: 8, value: 8},
        {key: 9, text: 9, value: 9},
        {key: 10, text: 10, value: 10},
        {key: 11, text: 11, value: 11},
        {key: 12, text: 12, value: 12},
        {key: 13, text: 13, value: 13},
        {key: 14, text: 14, value: 14},
        {key: 15, text: 15, value: 15},
        {key: 16, text: 16, value: 16},
        {key: 17, text: 17, value: 17},
        {key: 18, text: 18, value: 18},
        {key: 19, text: 19, value: 19},
        {key: 20, text: 20, value: 20},
        {key: 21, text: 21, value: 21},
        {key: 22, text: 22, value: 22},
        {key: 23, text: 23, value: 23},
        {key: 24, text: 24, value: 24},
        {key: 25, text: 25, value: 25},
        {key: 26, text: 26, value: 26},
        {key: 27, text: 27, value: 27},
        {key: 28, text: 28, value: 28},
        {key: 29, text: 29, value: 29},
        {key: 30, text: 30, value: 30},
        {key: 31, text: 31, value: 31},
    ]

    const cipherOptions = [
        {key: 1, text: 'Sin Cifrar', value: 1},
        {key: 2, text: 'Cifrado Sensible', value: 2},
        {key: 3, text: 'Cifrado', value: 3},
    ]

    const monthOptions = [
        {key: 1, text: 'ENERO', value: 1},
        {key: 2, text: 'FEBRERO', value: 2},
        {key: 3, text: 'MARZO', value: 3},
        {key: 4, text: 'ABRIL', value: 4},
        {key: 5, text: 'MAYO', value: 5},
        {key: 6, text: 'JUNIO', value: 6},
        {key: 7, text: 'JULIO', value: 7},
        {key: 8, text: 'AGOSTO', value: 8},
        {key: 9, text: 'SEPTIEMBRE', value: 9},
        {key: 10, text: 'OCTUBRE', value: 10},
        {key: 11, text: 'NOVIEMBRE', value: 11},
        {key: 12, text: 'DICIEMBRE', value: 12},
    ]

    const ArbolCuentas=()=>{
        async function obtener(){
            let res = await axios.get('http://localhost:3000/GetArbolCuentas/'+cipher)
            if (res.data.Tipo !== "Error"){
                setImagen("data:image/png;base64,"+res.data.Content)
                setTitle('Árbol Cuentas')
            }
        }
        obtener()
    }

    const Grafo = ()=>{
        async function obtener(){
            let res = await axios.get('http://localhost:3000/GetGrafo')
            if (res.data.Tipo !== "Error"){
                setImagen("data:image/png;base64,"+res.data.Content)
                setTitle('Grafo Completo')
            }
        }
        obtener()
    }

    const Vector = ()=>{
        async function obtener(){
            let res = await axios.get('http://localhost:3000/getArreglo')
            if (res.data.Tipo !== "Error"){
                setImagen("data:image/png;base64,"+res.data.Content)
                setTitle('Arreglo de Tiendas')
            }
        }
        obtener()
    }

    const arbolA = ()=>{
        async function obtener(){
            let res = await axios.get('http://localhost:3000/GetArbolAnio')
            if (res.data.Tipo !== "Error"){
                setImagen("data:image/png;base64,"+res.data.Content)
                setTitle('Árbol Años')
            }
        }
        obtener()
    }
    const arbolM = ()=>{
        async function obtener(){
            let res = await axios.get('http://localhost:3000/GetArbolMeses/'+year)
            if (res.data.Tipo !== "Error"){
                setImagen("data:image/png;base64,"+res.data.Content)
                setTitle('Árbol Meses')
            }
        }
        obtener()
    }
    const matriz = ()=>{
        async function obtener(){
            let res = await axios.get('http://localhost:3000/GetMatriz/'+year+'/'+month)
            if (res.data.Tipo !== "Error"){
                setImagen("data:image/png;base64,"+res.data.Content)
                setTitle('Matriz Pedidos')
            }
        }
        obtener()
    }
    const cola = ()=>{
        async function obtener(){
            let res = await axios.get('http://localhost:3000/GetPedidosDia/'+year+'/'+month+'/'+cat+'/'+day)
            if (res.data.Tipo !== "Error"){
                setImagen("data:image/png;base64,"+res.data.Content)
                setTitle('Cola de Pedidos')
            }
        }
        obtener()
    }
    return (
        <>
        <NavBar
        
        activo={0}
        />
        <div className="Content">
            <div className="ui segment mosaico container">
                <Segment>
                    <Header size="huge">
                        <Icon name='line graph'/>
                        <Header.Content>Reportes</Header.Content>
                    </Header>
                </Segment>
                <Container fluid>
                    <Grid>
                        <Grid.Row columns={3}>
                            <Grid.Column>
                                <Button color='teal' fluid onClick={Vector}>Obtener Arreglo<br/>de Tiendas</Button>
                            </Grid.Column>
                            <Grid.Column>
                                <Button color='teal' fluid onClick={arbolA}>Obtener Árbol<br/> Años</Button>
                            </Grid.Column>
                            <Grid.Column>
                                <Button color='teal' fluid onClick={Grafo}>Obtener Grafo<br/>Completo</Button>
                            </Grid.Column>
                        </Grid.Row>
                        <Grid.Row columns={2}>
                            <Grid.Column>
                                <Dropdown placeholder='Tipo...' selection fluid options={cipherOptions} onChange={ (e)=>{
                                    setCipher(getCipher(e.target.innerText))
                                }}/>
                            </Grid.Column>
                            <Grid.Column>
                                <center>
                                    <Button color='teal' onClick={ArbolCuentas}>Obtener Árbol de Cuentas</Button>
                                </center>
                            </Grid.Column>
                        </Grid.Row>
                        <Grid.Row columns={4}>
                            <Grid.Column>
                               <Input fluid placeholder='Año...' onChange={ (e)=>{
                                        setYear(e.target.value)
                                    }}/>
                            </Grid.Column>
                            <Grid.Column>
                               <Dropdown placeholder='Mes...' selection fluid options={monthOptions} onChange={ (e)=>{
                                        setMonth(getMonth(e.target.innerText))
                                    }}/> 
                            </Grid.Column>
                            <Grid.Column>
                               <Dropdown placeholder='Día...' selection fluid options={dayOptions} onChange={ (e)=>{
                                        setDay(e.target.innerText)
                                    }}/> 
                            </Grid.Column>
                            <Grid.Column>
                               <Input fluid placeholder='Categoria...'onChange={ (e)=>{
                                        setCat(e.target.value)
                                    }}/>
                            </Grid.Column>
                        </Grid.Row>
                        <Grid.Row columns={3}>
                            <Grid.Column>
                                <center>
                                    <Button color='teal' onClick={arbolM}>Obtener Árbol<br/>de Meses</Button>
                                </center>
                            </Grid.Column>
                            <Grid.Column>
                                <center>
                                    <Button color='teal' onClick={matriz}>Obtener Matriz<br/>de Pedidos</Button>
                                </center>
                            </Grid.Column>
                            <Grid.Column >
                                <center>
                                    <Button color='teal' onClick={cola}>Obtener Cola<br/>de Pedidos</Button>
                                </center>
                            </Grid.Column>
                        </Grid.Row>
                    </Grid>
                </Container>
                <Segment>
                    <Header size="huge">
                        <Header.Content>{title}</Header.Content>
                    </Header>
                </Segment>
                <Image fluid src={imagen}/>
            </div>
        </div>
       </> 
    )
}
export default Reporte
