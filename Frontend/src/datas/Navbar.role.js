const navbarContent = [
    {
        role: 'guest',
        content:[
            {title: 'Home', link: '/'},
            {title: 'Report', link: '/Report'},
            {title: 'Courses', link: '/Courses'},
            {title: 'Sign In', link: '/SignIn'},
            {title: 'Sign Up', link: '/SignUp'},
        ]
    },
    {
        role: 'tutor',
        content:[
            {title: 'Home', link: '/'},
            {title: 'My Courses', link: '/MyCourses'},
            {title: 'My Classes', link: '/MyClasses'},
            {title: 'Account', link: '/Account'},
            {title: 'Sign Out', link: '/SignOut'},
        ]
    },
    {
        role: 'student',
        content:[
            {title: 'Home', link: '/'},
            {title: 'My Classes', link: '/MyClasses'},
            {title: 'Account', link: '/Account'},
            {title: 'Sign Out', link: '/SignOut'},
        ]
    },
    {
        role: 'admin',
        content:[
            {title: 'Home', link: '/'},
            {title: 'Courses', link: '/Courses'},
            {title: 'Teachers', link: '/Teachers'},
            {title: 'Sign In', link: '/SignIn'},
            {title: 'Sign Up', link: '/SignUp'},
        ]
    }
];

export default navbarContent;