
interface NavbarOpen {
    open: boolean;
    setOpen: (value: boolean) => void;
}

function NavbarListbox({ open, setOpen }: NavbarOpen) {
    //const iconClasses = "text-xl text-default-500 pointer-events-none flex-shrink-0";
    const ListboxWrapper = `fixed top-0 left-0 w-full max-w-[260px] h-screen border-small px-1 py-2 rounded-small border-default-200 dark:border-default-100 z-50 backdrop-blur-md transform transition-transform duration-300 ${open ? "translate-x-0" : "-translate-x-full"}`;

    return (
        <div className={ListboxWrapper}>
      
        </div>
    );
}

export default NavbarListbox;
