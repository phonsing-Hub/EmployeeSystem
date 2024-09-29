import {
  Navbar,
  NavbarBrand,
  NavbarContent,
  NavbarItem,
} from "@nextui-org/react";
import DropdownUser from "./DropdownUser";

function NavbarHeader() {
  return (
    <Navbar shouldHideOnScroll maxWidth="2xl" isBordered>
      <NavbarBrand>
      <DropdownUser/>
      </NavbarBrand>
      <NavbarContent justify="end">
        <NavbarItem className="hidden lg:flex">
        <p className="font-bold text-inherit">ACME</p>
        </NavbarItem>
      </NavbarContent>
    </Navbar>
  );
}

export default NavbarHeader;
