import React, { useEffect, useState } from 'react';
import Table from '@material-ui/core/Table';
import TableHead from '@material-ui/core/TableHead';
import TableBody from '@material-ui/core/TableBody';
import TableRow from '@material-ui/core/TableRow';
import TableCell from '@material-ui/core/TableCell';
import axios from 'axios';

const StockList = () => {
    const [stocks, setStocks] = useState([])

    useEffect(() => {
        const getStocks = async () => {
            const res = await axios.get('http://localhost:3001/api/stocks')
            if (res) {
                console.log(res);
                setStocks(res.data)
            }
        }
        getStocks();
    }, []);

    return <Table>
        <TableHead>
            <TableRow>
                <TableCell> Symbol </TableCell>
                <TableCell> Price </TableCell>
                <TableCell> Trend </TableCell>
            </TableRow>
        </TableHead>
        <TableBody>
            {stocks.map((item: any) => <TableRow>
                <TableCell> {item.symbol} </TableCell>
                <TableCell> {item.price} </TableCell>
                <TableCell> {item.trend} </TableCell>
            </TableRow>)}
            
        </TableBody>
    </Table>    
}

export default StockList;
