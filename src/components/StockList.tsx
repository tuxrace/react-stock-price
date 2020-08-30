import React from 'react';
import Table from '@material-ui/core/Table';
import TableHead from '@material-ui/core/TableHead';
import TableBody from '@material-ui/core/TableBody';
import TableRow from '@material-ui/core/TableRow';
import TableCell from '@material-ui/core/TableCell';

const StockList = () => {
    return <Table>
        <TableHead>
            <TableRow>
                <TableCell> Symbol </TableCell>
                <TableCell> Price </TableCell>
                <TableCell> Trend </TableCell>
            </TableRow>
        </TableHead>
        <TableBody>
            <TableRow>
                <TableCell> D05 </TableCell>
                <TableCell> 20.1 </TableCell>
                <TableCell> Up </TableCell>
            </TableRow>
        </TableBody>
    </Table>    
}

export default StockList;
