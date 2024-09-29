import {SVGProps} from "react";

export type IconSvgProps = SVGProps<SVGSVGElement> & {
  size?: number;
};

export interface Users {
  id: number,
  firstname: string;
  lastname: string;
  email: string;
  phone: string;
  department: string;
  positions: string;
  salary: number;
}