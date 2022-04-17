import React from "react";
//import cells from "../data/cells";
import Cell from "./Cell";

export default class Area extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      cellsA: cells
    };
  }
  render() {
    return (
      <div>
        {this.state.cellsA.map((a, i) => {
          return a.map((v, j) => {
            return <Cell key={i + "n" + j} active={v} />;
          });
        })}
      </div>
    );
  }
}
