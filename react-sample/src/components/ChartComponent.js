import React, { useState } from "react";
import { Doughnut } from "react-chartjs-2";
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from "chart.js";

ChartJS.register(ArcElement, Tooltip, Legend);

const ChartComponent = () => {
  const [chartData, setChartData] = useState([30, 50, 20]);

  const data = {
    labels: ["Service entitled", "Non- service entitled", "None"],
    datasets: [
      {
        data: chartData,
        backgroundColor: ["#FF6384", "#36A2EB", "#FFCE56"],
      },
    ],
  };

  const totalValue = chartData.reduce((acc, val) => acc + val, 0);

  const updateData = () => {
    setChartData(chartData.map(() => Math.floor(Math.random() * 100)));
  };

  return (
    <div style={{ display: "flex", gap: "20px", alignItems: "center" }}>
      <div style={{ width: "300px", height: "300px" }}>
        <Doughnut data={data} />
      </div>
      <div style={{ textAlign: "center", padding: "20px", border: "1px solid gray" }}>
        <h2>Total Count</h2>
        <p style={{ fontSize: "36px", fontWeight: "bold", color: "#007bff" }}>{totalValue}</p>
        <button onClick={updateData} style={{ marginTop: "10px", padding: "5px 10px" }}>
          Update Data
        </button>
      </div>
    </div>
  );
};

export default ChartComponent;