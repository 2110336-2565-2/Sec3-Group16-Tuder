const navbarContent = [
    {
        role: 'guest',
        content:[
            {title: 'Home', link: '/'},
            {title: 'Sign In', link: '/sign-in'},
            {title: 'Sign Up', link: '/sign-up'},
        ]
    },
    {
        role: 'tutor',
        content:[
            {title: 'Home', link: '/'},
            {title: 'My Courses', link: '/my-courses'},
            {title: 'My Classes', link: '/my-classes'},
            {title: 'Account', dropdown:[
                {title: 'Profile', link: '/profile'},
                {title: 'Change Password', link: '/change-password'},
                {title: 'Sign Out', link: '/'},
            ]}
        ]
    },
    {
        role: 'student',
        content:[
            {title: 'Home', link: '/'},
            {title: 'My Classes', link: '/my-classes'},
            {title: 'Account', dropdown:[
                {title: 'Profile', link: '/profile'},
                {title: 'Change Password', link: '/change-password'},
                {title: 'Sign Out', link: '/'},
            ]}
        ]
    },
    {
        role: 'admin',
        content:[
            {title: 'Home', link: '/'},
            {title: 'Courses', link: '/courses'},
            {title: 'Teachers', link: '/teachers'},
            {title: 'Report List', link: '/issuereports'},
            {title: 'Cancel-Request List', link: '/cancel-request-list'},
            {title: 'Sign Out', link: '/'},
        ]
    }
];

export default navbarContent;