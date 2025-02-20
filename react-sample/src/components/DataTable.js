import React from "react";
import { useTable } from "@tanstack/react-table";

const DataTable = () => {
  const data = React.useMemo(
    () => [
      {
        name: "201580149",
        hostName: "-",
        make: "Cisco",
        model: "N9K-C93128TX",
        deviceType: "IP Switch",
        serialNumber: "SAL1949U995",
        country: "India",
        location: "MUMBAI, NARIMAN POINT, 400021",
        serviceEntitled: "No",
        impacted: "1",
        services: "-",
        actions: "View Details",
      },
    ],
    []
  );

  const columns = React.useMemo(
    () => [
      { Header: "Name", accessor: "name" },
      { Header: "Host Name", accessor: "hostName" },
      { Header: "Make", accessor: "make" },
      { Header: "Model", accessor: "model" },
      { Header: "Device Type", accessor: "deviceType" },
      { Header: "Serial Number", accessor: "serialNumber" },
      { Header: "Country", accessor: "country" },
      { Header: "Location", accessor: "location" },
      { Header: "Service Entitled", accessor: "serviceEntitled" },
      { Header: "Impacted", accessor: "impacted" },
      { Header: "Services", accessor: "services" },
      { Header: "Actions", accessor: "actions" },
    ],
    []
  );

  const { getTableProps, getTableBodyProps, headerGroups, rows, prepareRow } =
    useTable({ columns, data });

  return (
    <div style={{ width: "50%", padding: "20px" }}>
      <h3>Device Data Table</h3>
      <table {...getTableProps()} style={{ width: "100%", borderCollapse: "collapse", border: "1px solid #ddd" }}>
        <thead>
          {headerGroups.map((headerGroup) => (
            <tr {...headerGroup.getHeaderGroupProps()} style={{ background: "#f4f4f4", textAlign: "left" }}>
              {headerGroup.headers.map((column) => (
                <th {...column.getHeaderProps()} style={{ padding: "10px", borderBottom: "1px solid #ddd" }}>
                  {column.render("Header")}
                </th>
              ))}
            </tr>
          ))}
        </thead>
        <tbody {...getTableBodyProps()}>
          {rows.map((row) => {
            prepareRow(row);
            return (
              <tr {...row.getRowProps()} style={{ borderBottom: "1px solid #ddd" }}>
                {row.cells.map((cell) => (
                  <td {...cell.getCellProps()} style={{ padding: "10px" }}>
                    {cell.render("Cell")}
                  </td>
                ))}
              </tr>
            );
          })}
        </tbody>
      </table>
    </div>
  );
};

export default DataTable;