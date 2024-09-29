import { useNavigate } from "react-router-dom";
import { useUser } from "../../UserProvider";
import {
  Dropdown,
  DropdownTrigger,
  DropdownMenu,
  DropdownSection,
  DropdownItem,
  User,
} from "@nextui-org/react";
import { ThemeSwitcher } from "../theme/ThemeSwitcher";

export default function DropdownUser() {
  const navigate = useNavigate();
  const user = useUser();
  return (
    <Dropdown
      showArrow
      radius="sm"
      classNames={{
        base: "before:bg-default-200",
        content: "p-0 border-small border-divider bg-background",
      }}
    >
      <DropdownTrigger className="cursor-pointer">
        <User
          name={user?.firstname + " " + user?.lastname}
          description={user?.email}
          classNames={{
            name: "text-default-600 font-bold",
            description: "text-default-500",
          }}
          avatarProps={{
            size: "sm",
            //src: "https://avatars.githubusercontent.com/u/30373425?v=4",
            isBordered: true,
            color: "secondary",
          }}
        />
      </DropdownTrigger>
      <DropdownMenu
        aria-label="Custom item styles"
        disabledKeys={["profile"]}
        className="p-3"
        itemClasses={{
          base: [
            "rounded-md",
            "text-default-500",
            "transition-opacity",
            "data-[hover=true]:text-foreground",
            "data-[hover=true]:bg-default-100",
            "dark:data-[hover=true]:bg-default-50",
            "data-[selectable=true]:focus:bg-default-50",
            "data-[pressed=true]:opacity-70",
            "data-[focus-visible=true]:ring-default-500",
          ],
        }}
        onAction={(key) => navigate(key.toString())}
      >
        <DropdownSection aria-label="Profile & Actions" showDivider>
          <DropdownItem
            isReadOnly
            key="profile"
            className="opacity-100"
            textValue="Junior Garcia profile"
          >
            <User
               name={user?.firstname + " " + user?.firstname}
               description={user?.email}
              classNames={{
                name: "text-default-600 font-bold",
                description: "text-default-500",
              }}
              avatarProps={{
                size: "sm",
                //src: "https://avatars.githubusercontent.com/u/30373425?v=4",
                isBordered: true,
                color: "secondary",
              }}
            />
          </DropdownItem>
          <DropdownItem key="dashboard" textValue="Dashboard">
            Dashboard
          </DropdownItem>
          <DropdownItem key="employees" textValue="Employees">
            Employees
          </DropdownItem>
          <DropdownItem key="settings" textValue="Settings">
            Settings
          </DropdownItem>
        </DropdownSection>

        <DropdownSection aria-label="Preferences" showDivider>
          <DropdownItem key="quick_search" shortcut="⌘K" textValue="Quick search">
            Quick search
          </DropdownItem>
          <DropdownItem
            isReadOnly
            key="theme"
            className="cursor-default"
            textValue="Theme"
            endContent={<ThemeSwitcher />}
          >
            Theme
          </DropdownItem>
        </DropdownSection>

        <DropdownSection aria-label="Help & Feedback">
          <DropdownItem key="help_and_feedback" textValue="Help & Feedback">
            Help & Feedback
          </DropdownItem>
          <DropdownItem key="logout" textValue="Log Out">
            Log Out
          </DropdownItem>
        </DropdownSection>
      </DropdownMenu>
    </Dropdown>
  );
}
