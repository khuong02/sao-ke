import React, { CSSProperties } from "react";
import { Transaction } from "../types/Transaction";

interface SearchResultsProps {
  transactions: Transaction[];
}

const tableHeaderStyle: CSSProperties = {
  padding: "10px",
  borderBottom: "2px solid #ddd",
  backgroundColor: "#f4f4f4",
  textAlign: "left",
};

const tableCellStyle: CSSProperties = {
  padding: "10px",
  borderBottom: "1px solid #ddd",
};

const evenRowStyle: CSSProperties = {
  backgroundColor: "#fafafa",
};

const oddRowStyle: CSSProperties = {
  backgroundColor: "#f4f4f4",
};

const SearchResults: React.FC<SearchResultsProps> = ({ transactions }) => {
  return (
    <div style={{ overflowX: "auto" }}>
      <table
        style={{
          width: "100%",
          borderCollapse: "collapse",
          marginTop: "20px",
        }}
      >
        <thead>
          <tr>
            <th style={tableHeaderStyle}>Date</th>
            <th style={tableHeaderStyle}>Transaction No</th>
            <th style={tableHeaderStyle}>Credit</th>
            <th style={tableHeaderStyle}>Debit</th>
            <th style={tableHeaderStyle}>Detail</th>
          </tr>
        </thead>
        <tbody>
          {transactions.map((transaction, index) => (
            <tr
              key={index}
              style={index % 2 === 0 ? evenRowStyle : oddRowStyle}
            >
              <td style={tableCellStyle}>{transaction.date}</td>
              <td style={tableCellStyle}>{transaction.trans_no}</td>
              <td style={tableCellStyle}>{transaction.credit}</td>
              <td style={tableCellStyle}>{transaction.debit}</td>
              <td style={tableCellStyle}>{transaction.detail}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default SearchResults;
