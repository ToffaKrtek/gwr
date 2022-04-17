export default function Cell(props) {
  let classC = "cell";
  if (props.active) {
    classC += " active";
  }

  return <div className={classC}></div>;
}
