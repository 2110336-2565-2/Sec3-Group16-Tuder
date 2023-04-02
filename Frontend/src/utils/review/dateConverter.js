export default function formatDate(inputDate) {
    const date = new Date(inputDate);
    const day = date.getDate();
    const monthNames = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun',
                        'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'];
    const monthIndex = date.getMonth();
    const year = date.getFullYear();
  
    return `${day} ${monthNames[monthIndex]} ${year}`;
}