import styles from "./Spinner.module.css";
import cn from "classnames";

export function Spinner(): JSX.Element {
    return (
        <div className={cn(styles.spinner, styles.center)}>
            {[...Array(12)].map((_, i) => (
                <div key={i} className={styles["spinner-blade"]} />
            ))}
        </div>
    );
}
