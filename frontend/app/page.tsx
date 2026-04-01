"use client";

import { useEffect, useState } from "react";

type Skill = {
  id: number;
  name: string;
};

type Experience = {
  id: number;
  company: string;
  position: string;
  start_date: string;
  end_date: string;
  description: string;
};

type Education = {
  id: number;
  institution: string;
  degree: string;
  start_year: number;
  end_year: number;
};

type Volunteering = {
  id: number;
  organization: string;
  role: string;
  start_date: string;
  end_date: string;
  description: string;
};

type Interest = {
  id: number;
  name: string;
};

type CV = {
  name: string;
  title: string;
  email: string;
  phone: string;
  linkedin: string;
  about: string;
  skills: Skill[];
  experiences: Experience[];
  education: Education[];
  volunteering: Volunteering[];
  interests: Interest[];
};

export default function Home() {
  const [cv, setCv] = useState<CV | null>(null);

  useEffect(() => {
    fetch("http://localhost:8080/api/cv")
      .then((res) => res.json())
      .then((data) => setCv(data))
      .catch((err) => console.error(err));
  }, []);

  if (!cv) {
    return <div className="p-10">Loading...</div>;
  }

  const {
    name,
    title,
    email,
    phone,
    linkedin,
    about,
    skills,
    experiences,
    education,
    volunteering,
    interests,
  } = cv;

  return (
    <main className="min-h-screen bg-gray-50 p-10">

      {/* HERO */}
      <section className="text-center mb-20">
        <h1 className="text-5xl font-bold">{name}</h1>
        <p className="text-xl text-green-600 mt-2">{title}</p>
        <div className="mt-6 space-y-2">
          <p>{email}</p>
          <p>{phone}</p>
          <p>{linkedin}</p>
        </div>
      </section>

      {/* ABOUT */}
      <section className="max-w-3xl mx-auto mb-20 text-center">
        <h2 className="text-2xl font-bold mb-4">About</h2>
        <p>{about}</p>
      </section>

      {/* SKILLS */}
      <section className="max-w-4xl mx-auto mb-20">
        <h2 className="text-2xl font-bold mb-6 text-center">Skills</h2>
        <div className="flex flex-wrap gap-3 justify-center">
          {skills.map((skill) => (
            <span key={skill.id} className="px-4 py-2 bg-green-100 rounded-full text-sm">
              {skill.name}
            </span>
          ))}
        </div>
      </section>

      {/* EXPERIENCE */}
      <section className="max-w-4xl mx-auto mb-20">
        <h2 className="text-2xl font-bold mb-6 text-center">Experience</h2>
        <div className="space-y-8">
          {experiences.map((exp) => (
            <div key={exp.id} className="p-6 bg-white rounded-xl shadow">
              <h3 className="font-bold text-lg">{exp.position}</h3>
              <p className="text-sm text-gray-500">{exp.company}</p>
              <p className="text-sm text-gray-400">
                {exp.start_date} - {exp.end_date}
              </p>
              <p>{exp.description}</p>
            </div>
          ))}
        </div>
      </section>

      {/* EDUCATION */}
      <section className="max-w-4xl mx-auto mb-20">
        <h2 className="text-2xl font-bold mb-6 text-center">Education</h2>
        <div className="space-y-6">
          {education.map((edu) => (
            <div key={edu.id} className="p-6 bg-white rounded-xl shadow">
              <h3 className="font-bold">{edu.degree}</h3>
              <p className="text-sm text-gray-500">{edu.institution}</p>
              <p className="text-sm text-gray-400">
                {edu.start_year} - {edu.end_year}
              </p>
            </div>
          ))}
        </div>
      </section>

      {/* VOLUNTEERING */}
      <section className="max-w-4xl mx-auto mb-20">
        <h2 className="text-2xl font-bold mb-6 text-center">Volunteering</h2>
        <div className="space-y-6">
          {volunteering.map((vol) => (
            <div key={vol.id} className="p-6 bg-white rounded-xl shadow">
              <h3 className="font-bold">{vol.role}</h3>
              <p className="text-sm text-gray-500">{vol.organization}</p>
              <p className="text-sm text-gray-400">
                {vol.start_date} - {vol.end_date}
              </p>
              <p className="mt-2">{vol.description}</p>
            </div>
          ))}
        </div>
      </section>

      {/* INTERESTS */}
      <section className="max-w-4xl mx-auto mb-20">
        <h2 className="text-2xl font-bold mb-6 text-center">Interests</h2>
        <div className="flex flex-wrap gap-3 justify-center">
          {interests.map((interest) => (
            <span key={interest.id} className="px-4 py-2 bg-pink-100 rounded-full text-sm">
              {interest.name}
            </span>
          ))}
        </div>
      </section>

    </main>
  );
}
