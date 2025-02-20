import React from "react";
import ChartComponent from "./ChartComponent";
import DataTable from "./DataTable";

const Dashboard = () => {
  return (
    <div style={{ display: "flex", justifyContent: "space-between", padding: "20px" }}>
      <ChartComponent />
      <DataTable />
    </div>
  );
};

export default Dashboard;