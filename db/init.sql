CREATE TABLE IF NOT EXISTS cv (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    title TEXT NOT NULL,
    email TEXT,
    phone TEXT,
    linkedin TEXT,
    about TEXT
);

CREATE TABLE IF NOT EXISTS education (
    id SERIAL PRIMARY KEY,
    cv_id INT NOT NULL REFERENCES cv(id) ON DELETE CASCADE,
    institution TEXT NOT NULL,
    degree TEXT NOT NULL,
    start_year INT NOT NULL,
    end_year INT NOT NULL,
    description TEXT
);

CREATE TABLE IF NOT EXISTS volunteering (
    id SERIAL PRIMARY KEY,
    cv_id INT NOT NULL REFERENCES cv(id) ON DELETE CASCADE,
    organization TEXT NOT NULL,
    role TEXT NOT NULL,
    start_date TEXT NOT NULL,
    end_date TEXT NOT NULL,
    description TEXT
);

CREATE TABLE IF NOT EXISTS interests (
    id SERIAL PRIMARY KEY,
    cv_id INT NOT NULL REFERENCES cv(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    description TEXT
);

CREATE TABLE IF NOT EXISTS skills (
    id SERIAL PRIMARY KEY,
    cv_id INT NOT NULL REFERENCES cv(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    description TEXT
);

CREATE TABLE IF NOT EXISTS experience (
    id SERIAL PRIMARY KEY,
    cv_id INT NOT NULL REFERENCES cv(id) ON DELETE CASCADE,
    company TEXT NOT NULL,
    position TEXT NOT NULL,
    start_date TEXT NOT NULL,
    end_date TEXT NOT NULL,
    description TEXT
);

INSERT INTO cv (name, title, email, phone, linkedin, about)
VALUES (
    'Ezgi Ongan',
    'Computer Engineer',
    'oonganezgi@gmail.com',
    '+90 534 374 5319',
    'http://www.linkedin.com/in/ezgi-ongan-75b48b259',
    'My passion for computers began in childhood and grew during middle school while competing in the Future Problem Solving International, where we explored using technology for good. In high school, my love for math and aspiration to become an engineer led me to choose computer engineering. I continue advancing in cybersecurity professionally and personally, while also building expertise in data analysis and software development.'
);

INSERT INTO education (cv_id, institution, degree, start_year, end_year) VALUES
(1, 'Beşiktaş Atatürk Anatolian High School', 'High School Diploma', 2018, 2022),
(1, 'Bahçeşehir University', 'Bachelor of Science in Computer Engineering', 2022, 2026);

INSERT INTO volunteering (cv_id, organization, role, start_date, end_date, description) VALUES
(1, 'FPS Project International', 'Team Captain', '01.10.2016', '05.04.2017', 'Competed in the Future Problem Solving International competition, leading a team of students in solving complex technological challenges that will affect our lives in the future. Learned about futurizm and how to use technology for good, which inspired my passion for computer engineering and cybersecurity.'),
(1, 'Volleyball Club', 'Libero', '2014', 'Present', 'Volleyball has always been a part of my life. It was a passion that helped me develop discipline and teamwork skills.'),
(1, 'MUN Club', 'PR Head', '2020', '2022', 'Organized and managed public relations activities for the Model United Nations club, including coordinating with external partners and promoting events to increase student participation.'),
(1, 'Google Solution Challenge Project', 'Team Member', '2022', '2023', 'Collaborated on a project for the Google Solution Challenge, developing a solution to address a specific problem using Google technologies. Contributed to the design and implementation of the project, demonstrating teamwork and technical skills.'),
(1, 'Google Developer Student Clubs BAU', 'Social Media Leader', '2023', '2024', 'Led social media for the Google Developer Student Clubs at Bahçeşehir University, increasing engagement and awareness of club activities among students.'),
(1, 'Innovation & Management Club', 'Team Member', '2024', 'Present', 'Participated in the Innovation & Management Club, engaging in activities and discussions related to innovation, entrepreneurship, and management principles.');

INSERT INTO interests (cv_id, name) VALUES
(1, 'Cybersecurity'),
(1, 'Data Analysis'),
(1, 'Software Development'),
(1, 'Volleyball'),
(1, 'Windsurfing'),
(1, 'Traveling'),
(1, 'Cooking'),
(1, 'Learning Foreign Languages');

INSERT INTO skills (cv_id, name) VALUES
(1, 'Go'),
(1, 'PfSense'),
(1, 'Linux'),
(1, 'Management & Leadership'),
(1, 'Advanced Level C++'),
(1, 'MS Office Programmes'),
(1, 'PfSense'),
(1, 'Advanced Level English'),
(1, 'Beginner Level French'),
(1, 'ITIL 4'),
(1, 'SQL Queries'),
(1, 'Power BI'),
(1, 'Business Intelligence');

INSERT INTO experience (cv_id, company, position, start_date, end_date, description) VALUES
(1, 'Turcas Petrol & Shell', 'Long Term IT Intern', '09.03.2024', '15.03.2025', 'Resolved the requests submitted to the help desk application. (Manage Engine)Providing system and hardware support to my colleagues. Suggested and helped deliver a training program for company employees on the IT process. Ensured employees follow procedures properly, avoided off-the-books work, and sped up ticket resolution and as a result, IT operations efficiency increased by 40%.'),
(1, 'Eclit Bilgi Teknolojileri', 'Linux Team Intern', '18.03.2025', '18.06.2025', 'Monitored health and performance of 400+ Linux servers across 50+ client environments using Zabbix, achieving 99.8% uptime SLA compliance and reducing critical alert response time from 45 minutes to under 8 minutes.'),
(1, 'Eclit Bilgi Teknolojileri', 'Microsoft Team Intern', '18.06.2025', '18.09.2025', 'Daily monitored and assessed vulnerabilities on 250+ client Windows machines using custom SQL scripts in SQL Server Management Studio, identifying an average of 180 critical vulnerabilities per week and enabling patching within 24 hours.'),
(1, 'Eclit Bilgi Teknolojileri', 'Network Security Intern', '18.09.2025', '18.12.2025', 'Implemented configuration standards across all firewall appliances, streamlining network security management and reducing potential misconfiguration errors by 15% through standardized processes and documentation improvements.'),
(1, 'Eclit Bilgi Teknolojileri', 'Developer Team Intern', '18.12.2025', 'Present', 'Developed and implemented a webhook-based automation system that integrates with the company platform. Designed and built an ARP search mechanism to discover available IP addresses across VLANs and prefixes, automatically identifying unassigned addresses and registering them into the companys platform. Contributed to improving IP management efficiency and reducing manual network inventory processes.'),
(1, 'Eclit Bilgi Teknolojileri', 'CTO Assistancy Intern', '18.03.2025', 'Present', 'Built 5+ Power BI reports for project growth tracking, integrating 50K+ rows of sales data to monitor efficiency metrics, resulting in a 22% reduction in reporting time from 2 days to 4 hours per cycle.');